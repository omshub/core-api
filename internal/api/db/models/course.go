package models

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID           string         `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deletaed_at" gorm:"index"`
	Number       string         `json:"number"`
	Aliases      string         `json:"aliases"`
	Foundational bool           `json:"foundational,string"`
	Deprecated   bool           `json:"deprecated,string"`
	Link         string         `json:"link"`
	Department   string         `json:"department"`
	Name         string         `json:"name"`
	Reviews      []Review
}
