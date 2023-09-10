package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/khilmi-aminudin/todo-app/model"
	"github.com/khilmi-aminudin/todo-app/service"
)

type TodosHandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type todosHandler struct {
	todosService service.TodosService
}

func NewTodosHandler(todosService service.TodosService) TodosHandler {
	return &todosHandler{
		todosService: todosService,
	}
}

// CreateTodo godoc
// @Summary      create an todo
// @Description  create an todo for parent of items
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param		 todo	body		model.CreateTodoParam	true	"Add Todo to activity"
// @Success      201  {object}  model.SingleTodoResponse
// @Failure      400  {object}  model.HttpErrorResponse
// @Failure      404  {object}  model.HttpErrorResponse
// @Failure      500  {object}  model.HttpErrorResponse
// @Router       /todo-items [post]
// Create implements TodosHandler
func (h *todosHandler) Create(ctx *gin.Context) {
	var request model.Todos

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.HttpErrorResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	data, err := h.todosService.Create(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.HttpErrorResponse{
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, model.SingleTodoResponse{
		Status:  "Success Created",
		Message: "Success",
		Value:   data,
	})
}

// DeleteTodo godoc
// @Summary      Delete a todo
// @Description  Delete a todo for parent of items
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param		 id	path		int	true	"todo id to delete"
// @Success      200  {object}  model.WebResponse
// @Failure      400  {object}  model.HttpErrorResponse
// @Failure      404  {object}  model.HttpErrorResponse
// @Failure      500  {object}  model.HttpErrorResponse
// @Router       /todo-items/{id} [delete]
// Delete implements TodosHandler
func (h *todosHandler) Delete(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "Bad Request",
			Message: "Invalid id",
		})
		return
	}
	err = h.todosService.Delete(ctx, id)
	if err.Error() == "record not found" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Todo with ID %d Not Found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    model.EmptyStruct{},
	})
}

// GetTodo godoc
// @Summary      Get a todo
// @Description  Get a todo for parent of items
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param		 id	path	int	true	"todo id"
// @Success      200  {object}  model.SingleTodoResponse
// @Failure      400  {object}  model.HttpErrorResponse
// @Failure      404  {object}  model.HttpErrorResponse
// @Failure      500  {object}  model.HttpErrorResponse
// @Router       /todo-items/{id} [get]
// Get implements TodosHandler
func (h *todosHandler) Get(ctx *gin.Context) {
	paramID := ctx.Param("id")

	fmt.Println("paramID : ", paramID)
	id, err := strconv.Atoi(paramID)
	fmt.Println("ID : ", id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.HttpErrorResponse{
			Status:  "Bad Request",
			Message: "Invalid id",
		})
		return
	}

	data, err := h.todosService.Get(ctx, id)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, model.HttpErrorResponse{
				Status:  "Not Found",
				Message: fmt.Sprintf("Todo with ID %d Not Found", id),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, model.HttpErrorResponse{
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SingleTodoResponse{
		Status:  "Success",
		Message: "Success",
		Value:   data,
	})
}

// GetListTodos godoc
// @Summary      get all todos
// @Description  get list all of existing todos
// @Tags         todo
// @Produce      json
// @Param		 activity_group_id	path	int	false	"activity_group_id"
// @Success      200  {object}  model.ListTodosResponse
// @Failure      400  {object}  model.HttpErrorResponse
// @Failure      404  {object}  model.HttpErrorResponse
// @Failure      500  {object}  model.HttpErrorResponse
// @Router       /todo-items/{activity_group_id} [get]
// GetAll implements TodosHandler
func (h *todosHandler) GetAll(ctx *gin.Context) {
	stringActivityID := ctx.Query("activity_group_id")

	if stringActivityID != "" {
		actvityID, err := strconv.Atoi(stringActivityID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.HttpErrorResponse{
				Status:  "Bad Request",
				Message: "Invalid activity_group_id",
			})
			return
		}
		todos, err := h.todosService.GetAll(ctx, actvityID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.HttpErrorResponse{
				Status:  "Bad Request",
				Message: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, model.ListTodosResponse{
			Status:  "Success",
			Message: "Success",
			Value:   todos,
		})
		return
	}

	todos, err := h.todosService.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.HttpErrorResponse{
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.ListTodosResponse{
		Status:  "Success",
		Message: "Success",
		Value:   todos,
	})
}

// UpdateTodo godoc
// @Summary      update an todo
// @Description  update an todo
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param 		 id   path int true "id of todo"
// @Param		 todo	body		model.UpdateTodoParam	true	"Update Todo Param"
// @Success      200  {object}  model.SingleTodoResponse
// @Failure      400  {object}  model.HttpErrorResponse
// @Failure      404  {object}  model.HttpErrorResponse
// @Failure      500  {object}  model.HttpErrorResponse
// @Router       /todo-items/{id} [patch]
// Update implements TodosHandler
func (h *todosHandler) Update(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.HttpErrorResponse{
			Status:  "Bad Request",
			Message: "Invalid id",
		})
		return
	}

	var request model.Todos

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.HttpErrorResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	request.ID = id
	if err := h.todosService.Update(ctx, request); err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, model.HttpErrorResponse{
				Status:  "Not Found",
				Message: fmt.Sprintf("Todo with ID %d Not Found", id),
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, model.HttpErrorResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	todo, err := h.todosService.Get(ctx, request.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.HttpErrorResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SingleTodoResponse{
		Status:  "Success",
		Message: "Success",
		Value:   todo,
	})
}
