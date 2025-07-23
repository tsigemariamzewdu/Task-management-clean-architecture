package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"task_management/db"
	"task_management/Delivery/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func main() {
	err:= godotenv.Load()
	if err!=nil{
		
		log.Fatal("error loading env",err)
	}
	mongouri:=os.Getenv("mongo_url")
	clientOptions := options.Client().ApplyURI(mongouri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongoDB")
	db.TaskCollection = client.Database("taskmanager").Collection("tasks")
	db.UserCollection=client.Database("taskmanager").Collection("users")
	app := gin.Default()
	router.SetUpRoutes(app)
	app.Run("localhost:8081")

}
