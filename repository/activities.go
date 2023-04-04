package repository

import (
	"context"
	"time"

	"gorm.io/gorm"

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
	db *gorm.DB
}

func NewActivitiesRepository(db *gorm.DB) ActivitiesRepository {
	return &activitiesRepository{
		db: db,
	}
}

// Create implements ActivitiesRepository
func (r *activitiesRepository) Create(ctx context.Context, data model.Activities) (model.Activities, error) {
	if err := r.db.
		WithContext(ctx).
		Create(&data).
		Error; err != nil {
		return model.Activities{}, err
	}
	data.UpdatedAt, data.CreatedAt = time.Now(), time.Now()
	return data, nil
}

// Delete implements ActivitiesRepository
func (r *activitiesRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.
		WithContext(ctx).
		Delete(model.Activities{}, id).
		Error; err != nil {
		return err
	}
	return nil
}

// Get implements ActivitiesRepository
func (r *activitiesRepository) Get(ctx context.Context, id int) (model.Activities, error) {
	var m model.Activities
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return m, err
	}
	return m, nil
}

// GetAll implements ActivitiesRepository
func (r *activitiesRepository) GetAll(ctx context.Context) ([]model.Activities, error) {
	var m []model.Activities
	if err := r.db.WithContext(ctx).Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil

}

// Update implements ActivitiesRepository
func (r *activitiesRepository) Update(ctx context.Context, data model.Activities) error {
	if err := r.db.
		WithContext(ctx).
		Model(model.Activities{}).
		Where("id = ?", data.ID).
		Updates(data).
		Error; err != nil {
		return err
	}
	return nil
}
