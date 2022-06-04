// follow https://cgrant.medium.com/developing-a-simple-crud-api-with-go-gin-and-gorm-df87d98e6ed1

package handlers

import (
	"fmt"
	"net/http"
	"omshub/core-api/internal/api/db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAddProgramHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var program models.Program

		if err := c.BindJSON(&program); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			fmt.Println(err)
		}

		if result := db.Create(&program); result.Error != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(result.Error)
		}

		c.JSON(http.StatusOK, &program)
	}
}

func NewGetOneProgramHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var program models.Program
		if err := db.Where("id = ?", id).First(&program).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, program)
		}
	}
}

func NewGetAllProgramsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var programs []models.Program
		if err := db.Find(&programs).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, programs)
		}
	}
}

func NewUpdateProgramHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var program models.Program
		if err := db.Where("id = ?", id).First(&program).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			if err := c.BindJSON(&program); err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				fmt.Println(err)
			}
			db.Save(&program)
			c.JSON(http.StatusOK, program)
		}
	}
}

func NewDeleteProgramHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var program models.Program
		db.Where("id = ?", id).Delete(&program)
		c.JSON(http.StatusOK, program)
	}
}
