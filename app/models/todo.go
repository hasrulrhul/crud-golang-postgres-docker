package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          string         `gorm:"primary_key,type:uuid:" json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
