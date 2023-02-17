package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/khilmi-aminudin/todo-app/model"
	"github.com/khilmi-aminudin/todo-app/service"
)

type ActivitiesHandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type activitiesHandler struct {
	activitiesService service.ActivitiesService
}

func NewActivitiesHandler(activitiesService service.ActivitiesService) ActivitiesHandler {
	return &activitiesHandler{
		activitiesService: activitiesService,
	}
}

// Create implements ActivitiesHandler
func (h *activitiesHandler) Create(ctx *gin.Context) {
	var request model.Activities

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	data, err := h.activitiesService.Create(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, model.WebResponse{
		Status:  "Success",
		Message: "Success Created",
		Data:    data,
	})
}

// Delete implements ActivitiesHandler
func (h *activitiesHandler) Delete(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Invalid id",
		})
		return
	}

	if err := h.activitiesService.Delete(ctx, id); err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": fmt.Sprintf("Activity with ID %d Not Found", id),
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success Delete",
		// "data":
	})
}

// Get implements ActivitiesHandler
func (h *activitiesHandler) Get(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "Bad Request",
			Message: "Invalid id",
		})
		return
	}

	data, err := h.activitiesService.Get(ctx, id)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": fmt.Sprintf("Activity with ID %d Not Found", id),
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

// GetAll implements ActivitiesHandler
func (h *activitiesHandler) GetAll(ctx *gin.Context) {
	activities, err := h.activitiesService.GetAll(ctx)
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
		Data:    activities,
	})
}

// Update implements ActivitiesHandler
func (h *activitiesHandler) Update(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "Bad Request",
			Message: "Invalid id",
		})
		return
	}

	var request model.Activities

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	request.ID = id
	if err := h.activitiesService.Update(ctx, request); err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": fmt.Sprintf("Activity with ID %d Not Found", id),
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	activities, err := h.activitiesService.Get(ctx, request.ID)
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
		Data:    activities,
	})
}
