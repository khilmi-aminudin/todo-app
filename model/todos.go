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
	ActivityGroupID int        `json:"activity_group_id"`
	Activities      Activities `gorm:"foreignKey:ActivityGroupID" json:"-"`
	Title           string     `json:"title"`
	IsActive        *bool      `json:"is_active"`
	Priority        string     `json:"priority"`
	CreatedAt       time.Time  `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

var Query string
