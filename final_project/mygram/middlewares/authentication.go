package middlewares

import (
	"final_project/mygram/database"
	"final_project/mygram/helpers"
	"final_project/mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userId, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
				"error":    "Bad Request",
				"messsage": "Invalid Parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		UserID := uint(userData["id"].(float64))
		User := models.User{}

		err = db.Select("id").First(&User, uint(userId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doest'n exist",
			})
			return
		}

		if User.ID != UserID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
		}
		c.Next()
	}
}
