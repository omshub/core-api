// follow https://cgrant.medium.com/developing-a-simple-crud-api-with-go-gin-and-gorm-df87d98e6ed1

package handlers

import (
	"fmt"
	"net/http"
	"omshub/core-api/internal/api/db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAddReviewHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var review models.Review

		if err := c.BindJSON(&review); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			fmt.Println(err)
		}

		if result := db.Create(&review); result.Error != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(result.Error)
		}

		c.JSON(http.StatusOK, &review)
	}
}

func NewGetOneReviewHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var review models.ReviewAPI
		if err := db.Model(&models.Review{}).Where("id = ?", id).First(&review).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, review)
		}
	}
}

func NewGetAllReviewsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reviews []models.ReviewAPI
		if err := db.Model(&models.Review{}).Find(&reviews).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, reviews)
		}
	}
}

func NewUpdateReviewHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var review models.Review
		if err := db.Where("id = ?", id).First(&review).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			if err := c.BindJSON(&review); err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				fmt.Println(err)
			}
			db.Save(&review)
			c.JSON(http.StatusOK, review)
		}
	}
}

func NewDeleteReviewHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var review models.Review
		db.Where("id = ?", id).Delete(&review)
		c.JSON(http.StatusOK, review)
	}
}
