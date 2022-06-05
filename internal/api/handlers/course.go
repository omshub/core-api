// follow https://cgrant.medium.com/developing-a-simple-crud-api-with-go-gin-and-gorm-df87d98e6ed1

package handlers

import (
	"fmt"
	"net/http"
	"omshub/core-api/internal/api/db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAddCourseHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var course models.Course

		if err := c.BindJSON(&course); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			fmt.Println(err)
		}

		if result := db.Create(&course); result.Error != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(result.Error)
		}

		c.JSON(http.StatusOK, &course)
	}
}

func NewGetOneCourseHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var course models.Course
		if err := db.Where("id = ?", id).First(&course).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, course)
		}
	}
}

func NewGetAllCoursesHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var courses []models.Course
		if err := db.Find(&courses).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, courses)
		}
	}
}

func NewGetAllCourseReviewsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var reviews []models.Review
		if err := db.Where("course_id = ?", id).Find(&reviews).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			c.JSON(http.StatusOK, reviews)
		}
	}
}

func NewUpdateCourseHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var course models.Course
		if err := db.Where("id = ?", id).First(&course).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println(err)
		} else {
			if err := c.BindJSON(&course); err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				fmt.Println(err)
			}
			db.Save(&course)
			c.JSON(http.StatusOK, course)
		}
	}
}

func NewDeleteCourseHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var course models.Course
		db.Where("id = ?", id).Delete(&course)
		c.JSON(http.StatusOK, course)
	}
}
