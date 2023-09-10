package model

import (
	"time"
)

type Activities struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type CreateActivityParam struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type SingleActivityResponse struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Value   Activities `json:"value"`
}

type ListActivitiesResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Value   []Activities `json:"value,omitempty"`
}
