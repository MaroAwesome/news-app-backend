package controllers

import (
	"../models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//InsertNews inserts to DB
func InsertNews(news *models.News) *mongo.InsertOneResult {

	return models.InsertNews(news)
}

//GetAllNews retrives all news from DB
func GetAllNews() ([]models.News, error) {
	return models.GetAllNews()
}

// GetNew retrives an object in DB
func GetNew(id string) (models.News, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return *new(models.News), err
	}
	return models.GetNew(oid), nil
}

// UpdateNew updates an object in DB
func UpdateNew(new *models.News) *mongo.UpdateResult {
	return models.UpdateNew(new)
}
