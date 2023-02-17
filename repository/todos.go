package repository

import (
	"context"

	"gorm.io/gorm"

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
	db *gorm.DB
}

func NewTodosRepository(db *gorm.DB) TodosRepository {
	return &todosRepository{
		db: db,
	}
}

// Create implements TodosRepository
func (r *todosRepository) Create(ctx context.Context, data model.Todos) (int, error) {
	if err := r.db.WithContext(ctx).
		Table("todos").
		Create(&data).
		Error; err != nil {
		return 0, err
	}
	return data.ID, nil
}

// Delete implements TodosRepository
func (r *todosRepository) Delete(ctx context.Context, id int) error {
	var m model.Todos
	if err := r.db.WithContext(ctx).
		Table("todos").
		Where("id", id).
		Delete(&m).
		Error; err != nil {
		return err
	}
	return nil
}

// Get implements TodosRepository
func (r *todosRepository) Get(ctx context.Context, id int) (model.Todos, error) {
	var m model.Todos
	if err := r.db.WithContext(ctx).
		Table("todos").
		Where("id", id).
		First(&m).
		Error; err != nil {
		return m, err
	}
	return m, nil
}

// GetAll implements TodosRepository
func (r *todosRepository) GetAll(ctx context.Context, activityID ...int) ([]model.Todos, error) {
	var m []model.Todos

	if len(activityID) > 0 {
		if err := r.db.WithContext(ctx).
			Table("todos").
			Where("activity_group_id", activityID[0]).
			Scan(&m).
			Error; err != nil {
			return nil, err
		}

	} else {
		if err := r.db.WithContext(ctx).
			Table("todos").
			Scan(&m).
			Error; err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Update implements TodosRepository
func (r *todosRepository) Update(ctx context.Context, data model.Todos) error {
	if err := r.db.WithContext(ctx).
		Table("activities").
		Where("id", data.ID).
		Updates(&data).
		Error; err != nil {
		return err
	}
	return nil
}
