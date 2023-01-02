package database

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DatabaseInstance struct {
	Current       *mongo.Database
	Context       context.Context
	CancelContext context.CancelFunc
}

func Client() (*mongo.Client, error) {
	MongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		fmt.Println(err)
		return client, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		errorMsg := "MongoDB connection failed, connection timed-out"
		fmt.Println(errorMsg)
		return nil, errors.New(errorMsg)
	}

	if err != nil {
		fmt.Println(err)
		return client, err
	}
	fmt.Println("Connected to MongoDB!")

	return client, nil
}

var MongoMainInstance *mongo.Client

func Init() error {
	fmt.Println("Creating MongoDB connection...")
	client, err := Client()

	if err != nil {
		return err
	}

	MongoMainInstance = client
	return nil
}

func Database() DatabaseInstance {
	db := MongoMainInstance.Database(os.Getenv("MONGODB_DATABASE"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return DatabaseInstance{
		Current:       db,
		Context:       ctx,
		CancelContext: cancel,
	}
}
