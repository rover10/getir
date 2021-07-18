package main

import (
	"context"
	"fmt"
	"time"

	// import 'mongo-driver' package libraries

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Fields struct {
	CreatedAt string
	Counts    []int
	Key       string
}

func CreateNewMongoClient() (*mongo.Client, context.Context, error) {
	user := GetEnvVariableOrDefault("MONGO_USER", DEFAULT_MONGO_USER)
	password := GetEnvVariableOrDefault("MONGO_PASSWORD", DEFAULT_MONGO_PASSWORD)
	host := GetEnvVariableOrDefault("MONGO_HOST", DEFAULT_MONGO_HOST)
	database := GetEnvVariableOrDefault("MONGO_DATABASE", DEFAULT_MONGO_DATABASE)

	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true", user, password, host, database)
	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		//log.Error(err)
		return client, nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		//log.Error(err)
		return client, nil, err
	}
	return client, ctx, nil
}
