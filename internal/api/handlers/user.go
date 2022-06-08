// follow https://cgrant.medium.com/developing-a-simple-crud-api-with-go-gin-and-gorm-df87d98e6ed1

package handlers

import (
	"fmt"
	"net/http"
	"omshub/core-api/internal/api/db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAddUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			fmt.Println(err)
		}

		if result := db.Create(&user); result.Error != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(result.Error)
		}

		c.JSON(http.StatusOK, &user)
	}
}

func NewGetOneUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var user models.User
		if err := db.Where("id = ?", id).First(&user).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

func NewGetAllUsersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, users)
		}
	}
}

func NewUpdateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var user models.User
		if err := db.Where("id = ?", id).First(&user).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			if err := c.BindJSON(&user); err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				fmt.Println(err)
			}
			db.Save(&user)
			c.JSON(http.StatusOK, user)
		}
	}
}

func NewDeleteUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var user models.User
		db.Where("id = ?", id).Delete(&user)
		c.JSON(http.StatusOK, user)
	}
}
