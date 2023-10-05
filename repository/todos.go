package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/khilmi-aminudin/todo-app/model"
)

type TodosRepository interface {
	Create(ctx context.Context, data model.Todos) (model.Todos, error)
	Update(ctx context.Context, data model.Todos) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (model.Todos, error)
	GetAll(ctx context.Context, activityID ...int) ([]model.Todos, error)
}

type todosRepository struct {
	// db *gorm.DB
	db *sql.DB
}

func NewTodosRepository(db *sql.DB) TodosRepository {
	return &todosRepository{
		db: db,
	}
}

// Create implements TodosRepository
func (r *todosRepository) Create(ctx context.Context, data model.Todos) (model.Todos, error) {
	model.Query = "insert into todos (activity_group_id, title, is_active, priority) values (?,?,?,?);"
	res, err := r.db.ExecContext(ctx, model.Query, data.ActivityGroupID, data.Title, data.IsActive, data.Priority)
	if err != nil {
		return model.Todos{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return model.Todos{}, err
	}
	data.ID = int(id)
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	return data, nil
}

// Delete implements TodosRepository
func (r *todosRepository) Delete(ctx context.Context, id int) error {
	model.Query = "delete from todos where id = ?;"
	_, err := r.db.ExecContext(ctx, model.Query, id)
	if err != nil {
		return err
	}
	return nil
}

// Get implements TodosRepository
func (r *todosRepository) Get(ctx context.Context, id int) (model.Todos, error) {
	model.Query = "select id, activity_group_id, title, is_active, priority, created_at, updated_at from todos where id = ?;"
	row := r.db.QueryRowContext(ctx, model.Query, id)
	var m model.Todos
	if err := row.Scan(
		&m.ID,
		&m.ActivityGroupID,
		&m.Title,
		&m.IsActive,
		&m.Priority,
		&m.CreatedAt,
		&m.UpdatedAt,
	); err != nil {
		return m, err
	}
	return m, nil
}

// GetAll implements TodosRepository
func (r *todosRepository) GetAll(ctx context.Context, activityID ...int) ([]model.Todos, error) {
	var m []model.Todos
	var rows *sql.Rows
	var err error

	if len(activityID) > 0 {
		model.Query = "select id, activity_group_id, title, is_active, priority, created_at, updated_at from todos where activity_group_id = ?;"
		rows, err = r.db.QueryContext(ctx, model.Query, activityID[0])
	} else {
		model.Query = "select id, activity_group_id, title, is_active, priority, created_at, updated_at from todos;"
		rows, err = r.db.QueryContext(ctx, model.Query)
	}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var d model.Todos
		if err := rows.Scan(
			&d.ID,
			&d.ActivityGroupID,
			&d.Title,
			&d.IsActive,
			&d.Priority,
			&d.CreatedAt,
			&d.UpdatedAt,
		); err != nil {
			return nil, err
		}
		m = append(m, d)
	}

	return m, nil
}

// Update implements TodosRepository
func (r *todosRepository) Update(ctx context.Context, data model.Todos) error {
	model.Query = "update todos set activity_group_id = ?, title = ?, is_active = ?, priority = ? where id = ?;"
	_, err := r.db.ExecContext(ctx, model.Query, data.ActivityGroupID, data.Title, data.IsActive, data.Priority, data.ID)
	if err != nil {
		return err
	}
	return nil
}
