package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/khilmi-aminudin/todo-app/model"
	"github.com/khilmi-aminudin/todo-app/repository"
)

type ActivitiesService interface {
	Create(ctx context.Context, data model.Activities) (model.Activities, error)
	Update(ctx context.Context, data model.Activities) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (model.Activities, error)
	GetAll(ctx context.Context) ([]model.Activities, error)
}

type activitiesService struct {
	activitiesRespository repository.ActivitiesRepository
}

func NewActivitiesService(activitiesRespository repository.ActivitiesRepository) ActivitiesService {
	return &activitiesService{
		activitiesRespository: activitiesRespository,
	}
}

// Create implements ActivitiesService
func (s *activitiesService) Create(ctx context.Context, data model.Activities) (model.Activities, error) {
	if data.Title == "" {
		return model.Activities{}, errors.New("title cannot be null")
	}

	if len(data.Title) < 3 {
		return model.Activities{}, errors.New("title length must be greater than 2")
	}

	id, err := s.activitiesRespository.Create(ctx, data)
	if err != nil {
		return model.Activities{}, err
	}

	data.ID = id
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	return data, nil
}

// Delete implements ActivitiesService
func (s *activitiesService) Delete(ctx context.Context, id int) error {
	activity, err := s.activitiesRespository.Get(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("activity with id %d not found", id)
		}
		return err
	}

	if (activity == model.Activities{}) {
		return fmt.Errorf("activity with id %d not found", id)
	}

	if err := s.activitiesRespository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

// Get implements ActivitiesService
func (s *activitiesService) Get(ctx context.Context, id int) (model.Activities, error) {
	data, err := s.activitiesRespository.Get(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Activities{}, fmt.Errorf("activity with id %d not found", id)
		}
		return model.Activities{}, err
	}

	if (data == model.Activities{}) {
		return model.Activities{}, fmt.Errorf("activity with id %d not found", id)
	}

	return data, nil
}

// GetAll implements ActivitiesService
func (s *activitiesService) GetAll(ctx context.Context) ([]model.Activities, error) {
	activities, err := s.activitiesRespository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return activities, nil
}

// Update implements ActivitiesService
func (s *activitiesService) Update(ctx context.Context, data model.Activities) error {
	activity, err := s.activitiesRespository.Get(ctx, data.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("activity with id %d not found", data.ID)
		}
		return err
	}

	if (activity == model.Activities{}) {
		return fmt.Errorf("activity with id %d not found", data.ID)
	}

	if data.Title == "" && data.Email == "" {
		return errors.New("body cannot be empty")
	}

	if data.Title == "" {
		return errors.New("title cannt be empty")
	}

	if data.Email == "" {
		return errors.New("email cannt be empty")
	}

	data.UpdatedAt = time.Now()
	if err := s.activitiesRespository.Update(ctx, data); err != nil {
		return err
	}
	return nil
}
