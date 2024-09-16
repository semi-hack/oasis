package services

import (
	"context"

	"oasis/config"
	"oasis/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func FindCategories() ([]models.Category, error) {
	cursor, err := config.CategoryCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var categories []models.Category
	if err = cursor.All(context.TODO(), &categories); err != nil {
		return nil, err
	}
	return categories, nil
}

func CreateCategory(u models.Category) (*models.Category, error) {

	result, err := config.CategoryCollection.InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}
	u.ID = result.InsertedID.(primitive.ObjectID)
	return &u, nil
}