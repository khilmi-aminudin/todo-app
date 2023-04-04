package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/khilmi-aminudin/todo-app/model"
)

type ActivitiesRepository interface {
	Create(ctx context.Context, data model.Activities) (model.Activities, error)
	Update(ctx context.Context, data model.Activities) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (model.Activities, error)
	GetAll(ctx context.Context) ([]model.Activities, error)
}

type activitiesRepository struct {
	// db *gorm.DB
	db *sql.DB
}

//	func NewActivitiesRepository(db *gorm.DB) ActivitiesRepository {
//		return &activitiesRepository{
//			db: db,
//		}
//	}
func NewActivitiesRepository(db *sql.DB) ActivitiesRepository {
	return &activitiesRepository{
		db: db,
	}
}

// Create implements ActivitiesRepository
func (r *activitiesRepository) Create(ctx context.Context, data model.Activities) (model.Activities, error) {
	fmt.Println("TIME NOW :", data.CreatedAt)
	model.Query = "insert into activities (title, email) values (?,?);"
	res, err := r.db.ExecContext(ctx, model.Query, data.Title, data.Email)
	defer r.db.Close()
	if err != nil {
		return model.Activities{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return model.Activities{}, err
	}
	data.ID = int(id)
	return data, nil
}

// Delete implements ActivitiesRepository
func (r *activitiesRepository) Delete(ctx context.Context, id int) error {
	model.Query = "delete from activities where id = ?;"
	_, err := r.db.ExecContext(ctx, model.Query, id)
	defer r.db.Close()
	if err != nil {
		return err
	}
	return nil
}

// Get implements ActivitiesRepository
func (r *activitiesRepository) Get(ctx context.Context, id int) (model.Activities, error) {
	var m model.Activities
	model.Query = "select id, title, email, created_at, updated_at from activities where id = ?;"
	row := r.db.QueryRowContext(ctx, model.Query, id)
	defer r.db.Close()
	if err := row.Scan(
		&m.ID,
		&m.Title,
		&m.Email,
		&m.CreatedAt,
		&m.UpdatedAt,
	); err != nil {
		return m, err
	}

	return m, nil
}

// GetAll implements ActivitiesRepository
func (r *activitiesRepository) GetAll(ctx context.Context) ([]model.Activities, error) {
	var m []model.Activities
	model.Query = "select id, title, email, created_at, updated_at from activities;"
	rows, err := r.db.QueryContext(ctx, model.Query)
	defer r.db.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var d model.Activities
		if err := rows.Scan(
			&d.ID,
			&d.Title,
			&d.Email,
			&d.CreatedAt,
			&d.UpdatedAt,
		); err != nil {
			return nil, err
		}
		m = append(m, d)
	}
	return m, nil
}

// Update implements ActivitiesRepository
func (r *activitiesRepository) Update(ctx context.Context, data model.Activities) error {
	model.Query = "update activities set title = ?, email = ? where id = ?;"
	_, err := r.db.ExecContext(ctx, model.Query, data.Title, data.Email, data.ID)
	defer r.db.Close()
	if err != nil {
		return err
	}
	return nil
}
