package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//News is the news struct, model to be followed
type News struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Poster      string             `json:"poster,omitempty" bson:"poster,omitempty"`
	Type        string             `json:"type,omitempty" bson:"type,omitempty"`
	Date        time.Time          `json:"date,omitempty" bson:"date,omitempty"`
}

//InsertNews inserts to DB
func InsertNews(news *News) *mongo.InsertOneResult {

	collection := Client.Database("News").Collection("News")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, news)
	fmt.Println(result)
	return result
}

//GetAllNews retrives all news from DB
func GetAllNews() ([]News, error) {
	var news []News
	collection := Client.Database("News").Collection("News")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var new News
		cursor.Decode(&new)
		news = append(news, new)
	}
	return news, nil
}

// GetNew retrives an object in DB
func GetNew(id primitive.ObjectID) News {
	collection := Client.Database("News").Collection("News")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.D{{"_id", id}}
	new := News{}
	collection.FindOne(ctx, filter).Decode(&new)
	return new
}

// UpdateNew updates an object in DB
func UpdateNew(new *News) *mongo.UpdateResult {
	collection := Client.Database("News").Collection("News")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	old := GetNew(new.ID)
	result, err := collection.UpdateOne(ctx, old, bson.D{{"$set", new}})
	if err != nil {
		fmt.Errorf(err.Error())
	}
	return result
}
