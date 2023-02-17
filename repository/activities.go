package repository

import (
	"context"

	"gorm.io/gorm"

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
	db *gorm.DB
}

func NewActivitiesRepository(db *gorm.DB) ActivitiesRepository {
	return &activitiesRepository{
		db: db,
	}
}

// Create implements ActivitiesRepository
func (r *activitiesRepository) Create(ctx context.Context, data model.Activities) (int, error) {
	if err := r.db.WithContext(ctx).
		Table("activities").
		Create(&data).
		Error; err != nil {
		return 0, err
	}
	return data.ID, nil
}

// Delete implements ActivitiesRepository
func (r *activitiesRepository) Delete(ctx context.Context, id int) error {
	var m model.Activities
	if err := r.db.WithContext(ctx).
		Table("activities").
		Where("id", id).
		Delete(&m).
		Error; err != nil {
		return err
	}
	return nil
}

// Get implements ActivitiesRepository
func (r *activitiesRepository) Get(ctx context.Context, id int) (model.Activities, error) {
	var m model.Activities
	if err := r.db.WithContext(ctx).
		Table("activities").
		Where("id", id).
		First(&m).
		Error; err != nil {
		return m, err
	}
	return m, nil
}

// GetAll implements ActivitiesRepository
func (r *activitiesRepository) GetAll(ctx context.Context) ([]model.Activities, error) {
	var m []model.Activities
	if err := r.db.WithContext(ctx).
		Table("activities").
		Scan(&m).
		Error; err != nil {
		return nil, err
	}
	return m, nil
}

// Update implements ActivitiesRepository
func (r *activitiesRepository) Update(ctx context.Context, data model.Activities) error {
	if err := r.db.WithContext(ctx).
		Table("activities").
		Where("id", data.ID).
		Updates(&data).
		Error; err != nil {
		return err
	}
	return nil
}
