package controllers

import (
	"net/http"
	"fmt"

	"oasis/services"
	"oasis/models"
	"github.com/gin-gonic/gin"
)

func CreateRecommendation(c *gin.Context) {

	var req models.Recommendation
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req)



	user, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

	req.CreatorId = user.(*models.User).ID
	createdRecommendation, err := services.CreateRecommendation(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful", "data": createdRecommendation})
}

func GetRecommendations(c *gin.Context) {
	recommendations, err := services.FindRecommendations()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successful", "data": recommendations})
}

