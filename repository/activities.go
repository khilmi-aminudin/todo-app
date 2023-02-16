package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/khilmi-aminudin/todo-app/model"
)

type ActivitiesRepository interface {
	Create(ctx context.Context, data model.Activities) (int, error)
	Update(ctx context.Context, data model.Activities) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (model.Activities, error)
	GetAll(ctx context.Context) ([]model.Activities, error)
}

type activitiesRepository struct {
	db *sql.DB
}

func NewActivitiesRepository(db *sql.DB) ActivitiesRepository {
	return &activitiesRepository{
		db: db,
	}
}

var query string

// Create implements ActivitiesRepository
func (r *activitiesRepository) Create(ctx context.Context, data model.Activities) (int, error) {
	query = "INSERT INTO activities (title, email) VALUES (?,?)"

	res, err := r.db.ExecContext(ctx, query, data.Title, data.Email)
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

// Delete implements ActivitiesRepository
func (r *activitiesRepository) Delete(ctx context.Context, id int) error {
	query = "DELETE FROM activities WHERE activity_id = ?"
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

// Get implements ActivitiesRepository
func (r *activitiesRepository) Get(ctx context.Context, id int) (model.Activities, error) {
	query = "SELECT activity_id, title, email, created_at, updated_at FROM activities WHERE activity_id = ? LIMIT 1"
	row := r.db.QueryRowContext(ctx, query, id)

	var data model.Activities

	if err := row.Scan(
		&data.ID,
		&data.Title,
		&data.Email,
		&data.CreatedAt,
		&data.UpdatedAt,
	); err != nil {
		return model.Activities{}, err
	}
	return data, nil
}

// GetAll implements ActivitiesRepository
func (r *activitiesRepository) GetAll(ctx context.Context) ([]model.Activities, error) {
	var data []model.Activities
	query = "SELECT activity_id, title, email, created_at, updated_at FROM activities ORDER BY activity_id"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var d model.Activities
		rows.Scan(
			&d.ID,
			&d.Title,
			&d.Email,
			&d.CreatedAt,
			&d.UpdatedAt,
		)
		data = append(data, d)
	}

	return data, nil
}

// Update implements ActivitiesRepository
func (r *activitiesRepository) Update(ctx context.Context, data model.Activities) error {
	query = "UPDATE activities SET title = ?, email = ?, updated_at = ? WHERE activity_id = ?"

	res, err := r.db.ExecContext(ctx, query, data.Title, data.Email, data.UpdatedAt, data.ID)
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
