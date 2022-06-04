package models

import (
	"time"

	"gorm.io/gorm"
)

type Program struct {
	ID              string         `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name            string         `json:"name"`
	URL             string         `json:"url"`
	Specializations []Specialization
}
