package service

import (
	"context"
	"errors"

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

	res, err := s.activitiesRespository.Create(ctx, data)
	if err != nil {
		return model.Activities{}, err
	}
	return res, nil
}

// Delete implements ActivitiesService
func (s *activitiesService) Delete(ctx context.Context, id int) error {
	_, err := s.activitiesRespository.Get(ctx, id)
	if err != nil {
		return err
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
		return model.Activities{}, err
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
	_, err := s.activitiesRespository.Get(ctx, data.ID)
	if err != nil {
		return err
	}

	if data.Title == "" && data.Email == "" {
		return errors.New("title cannot be null")
	}

	if err := s.activitiesRespository.Update(ctx, data); err != nil {
		return err
	}
	return nil
}
