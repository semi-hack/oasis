package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)



var (
	ErrorInvalidEmail = "Invalid Email"
)

func Login(c *gin.Context) {
	login := &struct{
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"min=6"`
	}{}

	if err:= c.ShouldBindJSON(login); err != nil {
		c.JSON(http.StatusBadRequest, errors.New(ErrorInvalidEmail))
	}

	user, err := controllers.Login(login.Email, login.Password)

	claims := jwt.MapClaims{"user": user.GetID(), "exp": time.Now().Add(time.Hour*24)}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(env.JWT_SECRET))
	errors.CheckError(err)

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "login successful", "token": token, "data": user})
}