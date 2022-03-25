package models

import (
	"gorm.io/gorm"
	"time"
)

type Track struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Artist    string         `json:"artist"`
	Title     string         `json:"title"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
