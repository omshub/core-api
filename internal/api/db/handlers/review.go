// follow https://cgrant.medium.com/developing-a-simple-crud-api-with-go-gin-and-gorm-df87d98e6ed1

package handlers

import (
	"fmt"
	"net/http"
	"omshub/core-api/internal/api/db/models"

	"github.com/gin-gonic/gin"
)

func (h handler) AddReview(c *gin.Context) {
	var review models.Review

	if err := c.BindJSON(&review); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
	}

	if result := h.DB.Create(&review); result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(result.Error)
	}

	c.JSON(http.StatusOK, &review)
}

func (h handler) GetOneReview(c *gin.Context) {
	id := c.Params.ByName("id")
	var review models.Review
	if err := h.DB.Where("id = ?", id).First(&review).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, review)
	}
}

func (h handler) GetAllReviews(c *gin.Context) {
	var reviews []models.Review
	if err := h.DB.Find(&reviews).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, reviews)
	}
}

func (h handler) UpdateReview(c *gin.Context) {
	id := c.Params.ByName("id")
	var review models.Review
	if err := h.DB.Where("id = ?", id).First(&review).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		if err := c.BindJSON(&review); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			fmt.Println(err)
		}
		h.DB.Save(&review)
		c.JSON(http.StatusOK, review)
	}
}

func (h handler) DeleteReview(c *gin.Context) {
	id := c.Params.ByName("id")
	var review models.Review
	h.DB.Where("id = ?", id).Delete(&review)
	c.JSON(http.StatusOK, review)
}
