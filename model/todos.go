package model

import (
	"time"
)

const (
	PriorityVeryHigh string = "very-high"
	PriorityHigh     string = "high"
	PriorityNormal   string = "normal"
	PriorityLow      string = "low"
)

type Todos struct {
	ID              int        `json:"id" gorm:"primaryKey"`
	Title           string     `json:"title"`
	ActivityGroupID int        `json:"activity_group_id"`
	Activities      Activities `gorm:"foreignKey:ActivityGroupID" json:"-"`
	IsActive        bool       `json:"is_active" gorm:"default:true"`
	Priority        string     `json:"priority" gorm:"default:very-high"`
	CreatedAt       time.Time  `json:"created_at" gorm:"created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

var Query string

type CreateTodoParam struct {
	Title           string `json:"title"`
	ActivityGroupID int    `json:"activity_group_id" binding:"required"`
	Priority        string `json:"priority" gorm:"default:very-high"`
}

type UpdateTodoParam struct {
	Title           string `json:"title"`
	ActivityGroupID int    `json:"activity_group_id"`
	Priority        string `json:"priority"`
	IsActive        bool   `json:"is_active"`
}

type SingleTodoResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Value   Todos  `json:"value"`
}

type ListTodosResponse struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Value   []Todos `json:"value,omitempty"`
}
