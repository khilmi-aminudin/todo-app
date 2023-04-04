package repository

import (
	"context"
	"time"

	"gorm.io/gorm"

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
	db *gorm.DB
}

func NewTodosRepository(db *gorm.DB) TodosRepository {
	return &todosRepository{
		db: db,
	}
}

// Create implements TodosRepository
func (r *todosRepository) Create(ctx context.Context, data model.Todos) (model.Todos, error) {
	if err := r.db.
		WithContext(ctx).
		Create(&data).
		Error; err != nil {
		return model.Todos{}, err
	}
	data.UpdatedAt, data.CreatedAt = time.Now(), time.Now()
	return data, nil
}

// Delete implements TodosRepository
func (r *todosRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.
		WithContext(ctx).
		Delete(model.Todos{}, id).
		Error; err != nil {
		return err
	}
	return nil
}

// Get implements TodosRepository
func (r *todosRepository) Get(ctx context.Context, id int) (model.Todos, error) {
	var m model.Todos
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return m, err
	}
	return m, nil
}

// GetAll implements TodosRepository
func (r *todosRepository) GetAll(ctx context.Context, activityID ...int) ([]model.Todos, error) {
	var m []model.Todos

	db := r.db.WithContext(ctx)
	if len(activityID) > 0 {
		db = db.Where("activity_group_id = ?", activityID[0])
	}

	if err := db.Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// Update implements TodosRepository
func (r *todosRepository) Update(ctx context.Context, data model.Todos) error {
	if err := r.db.
		WithContext(ctx).
		Model(model.Todos{}).
		Where("id = ?", data.ID).
		Updates(data).
		Error; err != nil {
		return err
	}
	return nil
}
