package models

import (
	"time"
)

type BaseModel struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
}

type BlogPost struct {
	BaseModel
	Title   string `json:"title" gorm:"not null;size:1000"`
	Content string `json:"content" gorm:"not null;size:10000"`
}
