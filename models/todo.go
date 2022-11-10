package models

import "time"

type Todo struct {
	ID              int64     `json:"id"`
	ActivityGroupID int64     `json:"activity_group_id"`
	Title           string    `json:"title"`
	IsActive        string    `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

type TodoCreateParams struct {
	ActivityGroupID int64  `json:"activity_group_id" binding:"required"`
	Title           string `json:"title" binding:"required,min=3"`
	IsActive        string `json:"is_active" binding:"required"`
	Priority        string `json:"priority" binding:"required"`
}

type TodoUpdateParams struct {
	Title    string `json:"title" binding:"required,min=3"`
	IsActive string `json:"is_active" binding:"required"`
	Priority string `json:"priority" binding:"required"`
}
