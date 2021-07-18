package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *Server) PostInMemoryKeyVal(w http.ResponseWriter, r *http.Request) {
	body, err := RequestBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check header
	contentType := r.Header.Get(CONTENT_TYPE)
	if APPLICATION_JSON != contentType {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Request body not a JSON`))
		return
	}

	keyVal := ""
	valueVal := ""
	if key, ok := body["key"].(string); ok {
		keyVal = key
	} else {
		// Bad request
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"key" is missing`))
		return
	}

	if value, ok := body["value"].(string); ok {
		valueVal = value
	} else {
		// Bad request
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"value" is missing`))
		return
	}

	s.InmemoryDatabase[keyVal] = valueVal
	response := map[string]interface{}{"key": keyVal, "value": valueVal}
	resBytes, _ := json.Marshal(response)
	setJSONResponseHeader(w, http.StatusOK)
	w.Write(resBytes)
}

func (s *Server) GetInMemoryKey(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	key := values.Get("key")
	if key == "" {
		// No query param, return empty
		response := map[string]interface{}{}
		resBytes, _ := json.Marshal(response)
		setJSONResponseHeader(w, http.StatusOK)
		w.Write(resBytes)
		return
	}

	if value, ok := s.InmemoryDatabase[key]; ok {
		response := map[string]interface{}{"key": key, "value": value}
		resBytes, _ := json.Marshal(response)
		setJSONResponseHeader(w, http.StatusOK)
		w.Write(resBytes)
	}
}

func setJSONResponseHeader(w http.ResponseWriter, status int) {
	h := w.Header()
	h.Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) Mongo(w http.ResponseWriter, r *http.Request) {

	col := s.MongoClient.Database("getir-case-study").Collection("records")
	fmt.Println("Collection type:", reflect.TypeOf(col), "\n")

	var episodes []bson.M
	//cursor, err := col.Find(context.TODO(), bson.D{})
	cursor, err := col.Find(s.Context, &episodes)

	// Find() method raised an error
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(s.Context)

	} else {
		// iterate over docs using Next()
		for cursor.Next(s.Context) {
			// declare a result BSON object
			var result bson.M
			err := cursor.Decode(&result)
			// If there is a cursor.Decode error
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				//os.Exit(1)
				return
				// If there are no cursor.Decode errors
			} else {
				fmt.Println("\nresult type:", reflect.TypeOf(result))
				fmt.Println("result:", result)
			}
		}
	}
}
