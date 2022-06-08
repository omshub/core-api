// follow https://cgrant.medium.com/developing-a-simple-crud-api-with-go-gin-and-gorm-df87d98e6ed1

package handlers

import (
	"fmt"
	"net/http"
	"omshub/core-api/internal/api/db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAddSemesterHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var semester models.Semester

		if err := c.BindJSON(&semester); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			fmt.Println(err)
		}

		if result := db.Create(&semester); result.Error != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(result.Error)
		}

		c.JSON(http.StatusOK, &semester)
	}
}

func NewGetOneSemesterHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var semester models.Semester
		if err := db.Where("id = ?", id).First(&semester).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, semester)
		}
	}
}

func NewGetAllSemestersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var semesters []models.Semester
		if err := db.Find(&semesters).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, semesters)
		}
	}
}

func NewUpdateSemesterHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var semester models.Semester
		if err := db.Where("id = ?", id).First(&semester).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			if err := c.BindJSON(&semester); err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				fmt.Println(err)
			}
			db.Save(&semester)
			c.JSON(http.StatusOK, semester)
		}
	}
}

func NewDeleteSemesterHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var semester models.Semester
		db.Where("id = ?", id).Delete(&semester)
		c.JSON(http.StatusOK, semester)
	}
}
