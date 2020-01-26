package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var err error

//Init is responsible for initilizing connection with db
func Init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	clientOptions := options.Client().ApplyURI("mongodb://ec2-18-188-196-102.us-east-2.compute.amazonaws.com")
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Printf("Error: ", err)
	}
}
