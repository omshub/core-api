// follow https://cgrant.medium.com/developing-a-simple-crud-api-with-go-gin-and-gorm-df87d98e6ed1

package handlers

import (
	"fmt"
	"net/http"
	"omshub/core-api/internal/api/db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAddSpecializationHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var specialization models.Specialization

		if err := c.BindJSON(&specialization); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			fmt.Println(err)
		}

		if result := db.Create(&specialization); result.Error != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(result.Error)
		}

		c.JSON(http.StatusOK, &specialization)
	}
}

func NewGetOneSpecializationHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var specialization models.Specialization
		if err := db.Where("id = ?", id).First(&specialization).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, specialization)
		}
	}
}

func NewGetAllSpecializationsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var specializations []models.Specialization
		if err := db.Find(&specializations).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, specializations)
		}
	}
}

func NewUpdateSpecializationHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var specialization models.Specialization
		if err := db.Where("id = ?", id).First(&specialization).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			if err := c.BindJSON(&specialization); err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				fmt.Println(err)
			}
			db.Save(&specialization)
			c.JSON(http.StatusOK, specialization)
		}
	}
}

func NewDeleteSpecializationHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var specialization models.Specialization
		db.Where("id = ?", id).Delete(&specialization)
		c.JSON(http.StatusOK, specialization)
	}
}
