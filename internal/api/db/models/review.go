package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time      `json:"created_at"`
	CreatedAtLegacy int64          `json:"created,string"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Rating          uint           `json:"rating,string"`
	Difficulty      uint           `json:"difficulty,string"`
	Workload        uint           `json:"workload,string"`
	Body            string         `json:"body"`
	SemesterID      string         `json:"semester_id"`
	CourseID        string         `json:"course_id"`
	UserID          uint           `json:"user_id" gorm:"default:1"`
	Legacy          bool           `json:"is_legacy" gorm:"default:false"`
}

type ReviewAPI struct {
	CreatedAt  time.Time `json:"created_at"`
	SemesterID string    `json:"semester_id"`
	CourseID   string    `json:"course_id"`
	Rating     uint      `json:"rating"`
	Difficulty uint      `json:"difficulty"`
	Workload   uint      `json:"workload"`
	Body       string    `json:"body"`
	Legacy     bool      `json:"is_legacy" gorm:"default:false"`
}
