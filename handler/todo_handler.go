package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khilmi-aminudin/todo-app/models"
	"github.com/khilmi-aminudin/todo-app/service"
)

type TodoHandler interface {
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type todoHandler struct {
	service service.TodoService
}

func NewTodoHandler(service service.TodoService) TodoHandler {
	return &todoHandler{
		service: service,
	}
}

// Create implements TodoHandler
func (handler *todoHandler) Create(c *gin.Context) {
	var params models.TodoCreateParams

	err := c.ShouldBindJSON(&params)

	if err != nil {
		panic(err.Error())
	}

	response := handler.service.Create(params)

	c.JSON(http.StatusCreated, response)
}

// Delete implements TodoHandler
func (handler *todoHandler) Delete(c *gin.Context) {
	panic("unimplemented")
}

// Get implements TodoHandler
func (handler *todoHandler) Get(c *gin.Context) {
	panic("unimplemented")
}

// GetAll implements TodoHandler
func (handler *todoHandler) GetAll(c *gin.Context) {
	panic("unimplemented")
}

// Update implements TodoHandler
func (handler *todoHandler) Update(c *gin.Context) {
	panic("unimplemented")
}
