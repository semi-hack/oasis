package services

import (
	"context"
	"errors"
	"time"

	"oasis/config"
	"oasis/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindRecommendations() ([]models.User, error) {
	cursor, err := config.UserCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var users []models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

func CreateRecommendation(u models.Recommendation) (*models.Recommendation, error) {

	recommendation,_ := GetRecommendationByTitle(u.Title)
	if recommendation != nil {
		return nil, errors.New(ErrorUserAreadyExists)
	}

	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	result, err := config.RecommendationCollection.InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}
	u.ID = result.InsertedID.(primitive.ObjectID)
	return &u, nil
}

func GetRecommendationByTitle(title string) (*models.Recommendation, error) {
	var recommendation models.Recommendation
	err := config.RecommendationCollection.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&recommendation)
	if err != nil {
		return nil, err
	}

	return &recommendation, nil
}

