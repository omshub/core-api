package models

import (
	"time"

	"gorm.io/gorm"
)

type Semester struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Season    string         `json:"season"`
	Year      string         `json:"year"`
	Name      string         `json:"name"`
	Reviews   []Review
}
