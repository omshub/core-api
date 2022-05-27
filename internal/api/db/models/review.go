package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	CourseOfferingId int    `json:"course_offering_id"`
	OverallRating    int    `json:"overall_rating"`
	Difficulty       int    `json:"difficulty"`
	WorkloadHours    int    `json:"workload_hours"`
	Comment          string `json:"comment"`
}
