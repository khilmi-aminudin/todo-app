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
	IsActive        *bool      `json:"is_active"`
	Priority        string     `json:"priority"`
	CreatedAt       time.Time  `json:"createdAt" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time  `json:"updatedAt" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

var Query string
