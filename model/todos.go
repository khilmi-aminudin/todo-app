package model

import "time"

const (
	PriorityVeryHigh string = "very-high"
	PriorityHigh     string = "high"
	PriorityNormal   string = "normal"
	PriorityLow      string = "low"
)

type Todos struct {
	ID              int       `json:"id"`
	ActivityGroupID int       `json:"activity_group_id"`
	Title           string    `json:"title"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
