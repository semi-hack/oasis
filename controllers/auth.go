package controllers

import (
	"errors"
	"net/http"

	"oasis/services"
	"oasis/models"
	"github.com/gin-gonic/gin"
)

var (
	ErrorInvalidEmail = "Invalid Email"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := services.CreateUser(user)

	token, err := services.GenerateToken(user.Email)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusOK, gin.H{"message": "successful", "token": token, "user": createdUser})
}

func Login(c *gin.Context) {
	login := &struct{
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"min=6"`
	}{}

	if err:= c.ShouldBindJSON(login); err != nil {
		c.JSON(http.StatusBadRequest, errors.New(ErrorInvalidEmail))
	}

	
	zh, err := services.GetUserByEmail(login.Email)
	if err != nil {
		c.JSON(400, gin.H{"message": "error finding user"})
		return
	}

	hashedPassword := zh.Password
	password := login.Password

	err = services.CheckPasswordHash(password, hashedPassword)
	if err != nil {
		c.JSON(403, gin.H{"message": "Invalid user credentials"})
		return
	}
	
	if err != nil {
		c.JSON(403, gin.H{"message": "Invalid user credentials"})
		return
	}

	token, err := services.GenerateToken(login.Email)
	if err != nil {
		c.JSON(403, gin.H{"message": "There was a problem logging you in, try again later"})
		c.Abort()
		return
	}

	zh.Password = ""

	c.JSON(200, gin.H{"message": "success", "user": zh, "token": token })
}