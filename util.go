package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func RequestBody(r *http.Request) (map[string]interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return nil, err
	}

	log.Printf("RequestBody: \n%s", body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	m := make(map[string]interface{}, 0)
	json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func GetEnvVariableOrDefault(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}
	return val
}
