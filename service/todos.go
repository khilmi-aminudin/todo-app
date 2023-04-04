package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/khilmi-aminudin/todo-app/model"
	"github.com/khilmi-aminudin/todo-app/repository"
)

type TodosService interface {
	Create(ctx context.Context, data model.Todos) (model.Todos, error)
	Update(ctx context.Context, data model.Todos) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (model.Todos, error)
	GetAll(ctx context.Context, activityID ...int) ([]model.Todos, error)
}

type todosService struct {
	todosRespository repository.TodosRepository
}

func NewTodosService(todosRespository repository.TodosRepository) TodosService {
	return &todosService{
		todosRespository: todosRespository,
	}
}

// Create implements TodosService
func (s *todosService) Create(ctx context.Context, data model.Todos) (model.Todos, error) {
	if data.Title == "" {
		return model.Todos{}, errors.New("title cannot be null")
	}
	if data.ActivityGroupID <= 0 {
		return model.Todos{}, errors.New("activity_group_id cannot be null")
	}

	if data.Priority == "" {
		data.Priority = model.PriorityVeryHigh
	}

	todos, err := s.todosRespository.Create(ctx, data)

	if err != nil {
		return model.Todos{}, err
	}

	return todos, nil
}

// Delete implements TodosService
func (s *todosService) Delete(ctx context.Context, id int) error {
	_, err := s.todosRespository.Get(ctx, id)
	if err != nil {
		return err
	}

	if err := s.todosRespository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

// Get implements TodosService
func (s *todosService) Get(ctx context.Context, id int) (model.Todos, error) {
	data, err := s.todosRespository.Get(ctx, id)
	if err != nil {
		return model.Todos{}, err
	}

	if (data == model.Todos{}) {
		return model.Todos{}, fmt.Errorf("Todo with ID %d Not Found", id)
	}

	return data, nil

}

// GetAll implements TodosService
func (s *todosService) GetAll(ctx context.Context, activityID ...int) ([]model.Todos, error) {
	todos, err := s.todosRespository.GetAll(ctx, activityID...)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// Update implements TodosService
func (s *todosService) Update(ctx context.Context, data model.Todos) error {
	_, err := s.todosRespository.Get(ctx, data.ID)
	if err != nil {
		return err
	}

	if err := s.todosRespository.Update(ctx, data); err != nil {
		return err
	}
	return nil
}
