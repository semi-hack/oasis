package controllers

import (
	"net/http"

	"oasis/services"
	"oasis/models"
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var req models.Category
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCategory, err := services.CreateCategory(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful", "data": createdCategory})
}

func GetCategory(c *gin.Context) {
	categories, err := services.FindCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successful", "data": categories})
}
