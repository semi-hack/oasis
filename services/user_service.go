package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"oasis/config"
	"oasis/models"
	"oasis/utils/validators"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)


var (
	ErrorFailedToFetchRecord = "failed to fetch record"
	ErrorFailedToUnMarshallRecord = "failed to UnMarshall record"
	ErrorInvalidEmail = "Invalid Email"
	ErrorUserAreadyExists = "This user already exists"
	ErrorCouldNotMarshallItem = "could not marshall item"
	ErrorUserDoesNotExist = "User does not exist"
	ErrorCouldNotDeleteItem = "Could not delete item"
)


func FindUsers() ([]models.User, error) {
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

func CreateUser(u models.User) (*models.User, error) {

	if !validators.IsEmailValid(u.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}

	currentUser,_ := GetUserByEmail(u.Email)
	if currentUser != nil {
		return nil, errors.New(ErrorUserAreadyExists)
	}

	hashed, _ := HashPassword(u.Password)

	u.Password = string(hashed)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	result, err := config.UserCollection.InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}
	u.ID = result.InsertedID.(primitive.ObjectID)
	u.Password = ""
	fmt.Println(u)
	return &u, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.UserCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password")
	}
	return nil
}