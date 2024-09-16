package routes

import (
	"oasis/services"
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrorAuth = "authentication error"
)

func Handle404(c *gin.Context) {
	c.AbortWithStatusJSON(404, gin.H{
		"success": false,
		"message": "API Endpoint not found",
		"error":   "INVALID_ROUTE_ERROR",
	})
}

func Handle405(c *gin.Context) {
	c.AbortWithStatusJSON(405, gin.H{
		"success": false,
		"message": "Method not allowed",
		"error":   "INVALID_METHOD_ERROR",
	})
}

func UserValidation(c *gin.Context) {
	if len(c.GetHeader("authorization")) < 7 {
		c.AbortWithStatusJSON(400, gin.H{"success": false, "message": "No auth token", "error": errors.New(ErrorAuth)})
		return
	}
	token := c.GetHeader("authorization")[7:]

	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{"success": false, "message": "Please login to continue.", "error": errors.New(ErrorAuth)})
		return
	}

	claims, err := services.ValidateToken(token)
	if err != nil {
		panic(err)
	}

	// Extract the email from claims and set it in the context
	if email, ok := claims["email"].(string); ok {
		user, err := services.GetUserByEmail(email)
		if err != nil {
			panic(err)
		}
		c.Set("user", user)
	} else {
		c.JSON(401, gin.H{"error": "email claim not found"})
		c.Abort()
		return
	}

	// Continue to the next handler
	c.Next()
}