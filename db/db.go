package db

import (
	"context"
	"fmt"
	
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	username = "username"
	password = "password"
	database = "db"
)

var (
	client     *mongo.Client
	clientOnce sync.Once
)

func getClient() (*mongo.Client, error) {
	var err error
	clientOnce.Do(func() {
		uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.isgee.mongodb.net/", username, password)
		clientOptions := options.Client().ApplyURI(uri)
		
		client, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			return
		}
		
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			return
		}
		
		fmt.Println("Connected to MongoDB!")
	})
	
	return client, err
}

func GetUsersCollection() (*mongo.Collection) {
	client, err := getClient()
	if err != nil {
		return nil
	}
	return client.Database(database).Collection("users")
}

func GetTasksCollection() (*mongo.Collection) {
	client, err := getClient()
	if err != nil {
		return nil
	}
	return client.Database(database).Collection("tasks")
}