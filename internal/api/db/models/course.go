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

type CourseAPI struct {
	ID           string `json:"id"`
	Department   string `json:"department"`
	Number       string `json:"number"`
	Name         string `json:"name"`
	Aliases      string `json:"aliases"`
	Foundational bool   `json:"foundational"`
	Deprecated   bool   `json:"deprecated"`
	Link         string `json:"link"`
}

type CourseStatAPI struct {
	ID            string  `json:"id"`
	AvgRating     float32 `json:"avg_rating"`
	AvgDifficulty float32 `json:"avg_difficulty"`
	AvgWorkload   float32 `json:"avg_workload"`
}
