package repository

import (
	"database/sql"

	"github.com/khilmi-aminudin/todo-app/models"
)

type TodoRepository interface {
	Get(todo models.Todo) models.Todo
	GetAll() []models.Todo
	Create(todo models.Todo) (models.Todo, error)
	Update(todo models.Todo) models.Todo
	Delete(todo models.Todo) error
}

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

// activity_group_id, title, is_active, priority

// Create implements TodoRepository
func (repo *todoRepository) Create(todo models.Todo) (models.Todo, error) {
	queryString = `
	INSERT INTO todo (
		activity_group_id,
		title,
		is_active, 
		priority
	) VALUES (
		$1, $2, $3, $4
	) RETURNING
	id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at;
	`

	row := repo.db.QueryRow(queryString, todo.ActivityGroupID, todo.Title, todo.IsActive, todo.Priority)

	var t models.Todo

	err := row.Scan(
		&t.ID,
		&t.ActivityGroupID,
		&t.Title,
		&t.IsActive,
		&t.Priority,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.DeletedAt,
	)

	return t, err
}

// Delete implements TodoRepository
func (repo *todoRepository) Delete(todo models.Todo) error {
	panic("unimplemented")
}

// Get implements TodoRepository
func (repo *todoRepository) Get(todo models.Todo) models.Todo {
	panic("unimplemented")
}

// GetAll implements TodoRepository
func (repo *todoRepository) GetAll() []models.Todo {
	panic("unimplemented")
}

// Update implements TodoRepository
func (repo *todoRepository) Update(todo models.Todo) models.Todo {
	panic("unimplemented")
}
