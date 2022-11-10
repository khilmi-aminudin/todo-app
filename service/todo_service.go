package service

import (
	"fmt"

	"github.com/khilmi-aminudin/todo-app/models"
	"github.com/khilmi-aminudin/todo-app/repository"
	"github.com/khilmi-aminudin/todo-app/web"
)

type TodoService interface {
	Get(id int64) web.WebResponse
	GetAll() web.WebResponse
	Create(todoCreateParams models.TodoCreateParams) web.WebResponse
	Update(activity web.ActivityUpdateRequest) web.WebResponse
	Delete(id int64) web.WebResponse
}

type todoService struct {
	repository repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{
		repository: repo,
	}
}

// Create implements TodoService
func (service *todoService) Create(todoCreateParams models.TodoCreateParams) web.WebResponse {
	todo := models.Todo{
		ActivityGroupID: todoCreateParams.ActivityGroupID,
		Title:           todoCreateParams.Title,
		IsActive:        todoCreateParams.IsActive,
		Priority:        todoCreateParams.Priority,
	}

	data, err := service.repository.Create(todo)

	if err != nil {
		fmt.Println(err.Error())
		return web.WebResponse{
			Status:  "Bad Request",
			Message: "Have an error when create todo",
		}
	}

	return web.WebResponse{
		Status:  "Success",
		Message: "todo item created",
		Data:    data,
	}
}

// Delete implements TodoService
func (service *todoService) Delete(id int64) web.WebResponse {
	panic("unimplemented")
}

// Get implements TodoService
func (service *todoService) Get(id int64) web.WebResponse {
	panic("unimplemented")
}

// GetAll implements TodoService
func (service *todoService) GetAll() web.WebResponse {
	panic("unimplemented")
}

// Update implements TodoService
func (service *todoService) Update(activity web.ActivityUpdateRequest) web.WebResponse {
	panic("unimplemented")
}
