package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/khilmi-aminudin/todo-app/model"
)

type TodosRepository interface {
	Create(ctx context.Context, data model.Todos) (int, error)
	Update(ctx context.Context, data model.Todos) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (model.Todos, error)
	GetAll(ctx context.Context, activityID ...int) ([]model.Todos, error)
}

type todosRepository struct {
	db *sql.DB
}

func NewTodosRepository(db *sql.DB) TodosRepository {
	return &todosRepository{
		db: db,
	}
}

// Create implements TodosRepository
func (r *todosRepository) Create(ctx context.Context, data model.Todos) (int, error) {
	query = "INSERT INTO todos (activity_group_id, title, priority, is_active) VALUES (?,?,?,?)"

	res, err := r.db.ExecContext(ctx, query, data.ActivityGroupID, data.Title, data.IsActive, data.Priority)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Delete implements TodosRepository
func (r *todosRepository) Delete(ctx context.Context, id int) error {
	query = "DELETE FROM todos WHERE todo_id = ?"
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if aff < 1 {
		return errors.New("delete failed")
	}
	return nil
}

// Get implements TodosRepository
func (r *todosRepository) Get(ctx context.Context, id int) (model.Todos, error) {
	query = "SELECT todo_id, activity_group_id, title, priority, is_active, created_at, updated_at FROM todos WHERE id = ? LIMIT 1"
	row := r.db.QueryRowContext(ctx, query, id)

	var data model.Todos

	if err := row.Scan(
		&data.ID,
		&data.ActivityGroupID,
		&data.Title,
		&data.Priority,
		&data.IsActive,
		&data.CreatedAt,
		&data.UpdatedAt,
	); err != nil {
		return model.Todos{}, err
	}
	return data, nil
}

// GetAll implements TodosRepository
func (r *todosRepository) GetAll(ctx context.Context, activityID ...int) ([]model.Todos, error) {
	var data []model.Todos

	if len(activityID) > 0 {
		query = "SELECT todo_id, activity_group_id, title, priority, is_active, created_at, updated_at  FROM todos WHERE activity_group_id = ? ORDER BY id"

		rows, err := r.db.QueryContext(ctx, query, activityID[0])
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var d model.Todos
			rows.Scan(
				&d.ID,
				&d.ActivityGroupID,
				&d.Title,
				&d.Priority,
				&d.IsActive,
				&d.CreatedAt,
				&d.UpdatedAt,
			)
			data = append(data, d)
		}

	} else {
		query = "SELECT todo_id, activity_group_id, title, priority, is_active, created_at, updated_at FROM todos ORDER BY id"
		rows, err := r.db.QueryContext(ctx, query)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var d model.Todos
			rows.Scan(
				&d.ID,
				&d.ActivityGroupID,
				&d.Title,
				&d.Priority,
				&d.IsActive,
				&d.CreatedAt,
				&d.UpdatedAt,
			)
			data = append(data, d)
		}
	}
	return data, nil
}

// Update implements TodosRepository
func (r *todosRepository) Update(ctx context.Context, data model.Todos) error {
	query = "UPDATE todos SET activity_group_id = ?, title = ?, is_active = ?, priority = ?, updated_at = ? WHERE todo_id = ?"

	res, err := r.db.ExecContext(ctx, query, data.ActivityGroupID, data.Title, data.IsActive, data.Priority, data.UpdatedAt, data.ID)
	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if aff < 1 {
		return errors.New("update failed")
	}
	return nil
}
