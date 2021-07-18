package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	InmemoryDatabase map[string]string
	MongoClient      *mongo.Client
	Context          context.Context
}

func main() {
	log.Printf("creating mongo client")
	mongoClient, ctx, err := CreateNewMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("mongo client created")
	defer mongoClient.Disconnect(ctx)

	log.Printf("setting up server")
	server := Server{
		InmemoryDatabase: make(map[string]string, 0),
		MongoClient:      mongoClient,
		Context:          ctx,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server.Mongo(w, r)
	})

	http.HandleFunc("/in-memory", func(w http.ResponseWriter, r *http.Request) {
		if GET == r.Method {
			server.GetInMemoryKey(w, r)
		} else if POST == r.Method {
			server.PostInMemoryKeyVal(w, r)
		}
	})

	port := GetEnvVariableOrDefault("PORT", DEFAULT_PORT)
	host := fmt.Sprintf("%s:%s", "localhost", port)
	log.Printf("server setup complete")
	log.Printf("running on : %s", host)
	log.Fatal(http.ListenAndServe(host, nil))

}
