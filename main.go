package main

import (
	"context"
	"time"

	"github.com/corey888773/ztp-api/srv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://ztp-mongo:27017"))
	if err != nil {
		panic(err)
	}

	server := srv.NewServer()
	server.SetupRouter()
	server.SetupDatabase(mongoClient)
	err = server.Start(":8000")
	if err != nil {
		panic(err)
	}
}
