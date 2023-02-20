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

// Create implements TodosHandler
func (h *todosHandler) Create(ctx *gin.Context) {
	var request model.Todos

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	data, err := h.todosService.Create(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, model.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	})
}

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
	if err := h.todosService.Delete(ctx, id); err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": fmt.Sprintf("Todo with ID %d Not Found", id),
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    model.EmptyStruct{},
	})
}

// Get implements TodosHandler
func (h *todosHandler) Get(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "Bad Request",
			Message: "Invalid id",
		})
		return
	}

	data, err := h.todosService.Get(ctx, id)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": fmt.Sprintf("Todo with ID %d Not Found", id),
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	})
}

// GetAll implements TodosHandler
func (h *todosHandler) GetAll(ctx *gin.Context) {
	stringActivityID := ctx.Query("activity_group_id")

	if stringActivityID != "" {
		actvityID, err := strconv.Atoi(stringActivityID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.WebResponse{
				Status:  "Bad Request",
				Message: "Invalid activity_group_id",
			})
			return
		}
		todos, err := h.todosService.GetAll(ctx, actvityID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.WebResponse{
				Status:  "Bad Request",
				Message: err.Error(),
			})
			return
		}

		if todos != nil {
			ctx.JSON(http.StatusOK, model.WebResponse{
				Status:  "Success",
				Message: "Success",
				Data:    todos,
			})
			return
		}
		ctx.JSON(http.StatusOK, model.WebResponse{
			Status:  "Success",
			Message: "Success",
			Data:    []model.EmptyStruct{},
		})
		return
	}

	todos, err := h.todosService.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	if todos != nil {
		ctx.JSON(http.StatusOK, model.WebResponse{
			Status:  "Success",
			Message: "Success",
			Data:    todos,
		})
		return
	}

	ctx.JSON(http.StatusOK, model.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    []model.EmptyStruct{},
	})
}

// Update implements TodosHandler
func (h *todosHandler) Update(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "Bad Request",
			Message: "Invalid id",
		})
		return
	}

	var request model.Todos

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	request.ID = id
	if err := h.todosService.Update(ctx, request); err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": fmt.Sprintf("Activity with ID %d Not Found", id),
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	todo, err := h.todosService.Get(ctx, request.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	})
}
