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

// CreateActivity godoc
// @Summary      create an activity
// @Description  create an activity for parent of items
// @Tags         activity
// @Accept       json
// @Produce      json
// @Param		 activity	body		model.CreateActivityParam	true	"Add Activities"
// @Success      201  {object}  model.SingleActivityResponse
// @Failure      400  {object}  model.HttpErrorResponse
// @Failure      404  {object}  model.HttpErrorResponse
// @Failure      500  {object}  model.HttpErrorResponse
// @Security BearerToken
// @Router       /activity-groups [post]
func (h *activitiesHandler) Create(ctx *gin.Context) {
	var request model.CreateActivityParam

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.HttpErrorResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	var data = model.Activities{
		Email: request.Email,
		Title: request.Title,
	}

	data, err := h.activitiesService.Create(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.HttpErrorResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, model.SingleActivityResponse{
		Status:  "Success Created",
		Message: "Success",
		Value:   data,
	})
}

// DeleteActivity godoc
// @Summary      delete a activity
// @Description  delete single activity
// @Tags         activity
// @Produce      json
// @Param 		 id   path int true "id of activity"
// @Success      200  {object}  model.WebResponse
// @Failure      400  {object}  model.HttpErrorResponse
// @Failure      404  {object}  model.HttpErrorResponse
// @Failure      500  {object}  model.HttpErrorResponse
// @Router       /activity-groups/{id} [delete]
// @Security BearerToken
// Delete implements ActivitiesHandler
func (h *activitiesHandler) Delete(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.HttpErrorResponse{
			Status:  "Bad Request",
			Message: "Invalid id",
		})
		return
	}

	err = h.activitiesService.Delete(ctx, id)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, model.HttpErrorResponse{
				Status:  "Not Found",
				Message: fmt.Sprintf("Activity with ID %d Not Found", id),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, model.HttpErrorResponse{
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.WebResponse{
		Status:  "Success",
		Message: "Success",
	})
}

// GetActivity godoc
// @Summary      get a activity
// @Description  get single activity
// @Tags         activity
// @Produce      json
// @Param 		 id   path int true "id of activity"
// @Success      200  {object}  model.SingleActivityResponse
// @Failure      400  {object}  model.HttpErrorResponse
// @Failure      404  {object}  model.HttpErrorResponse
// @Failure      500  {object}  model.HttpErrorResponse
// @Router       /activity-groups/{id} [get]
// @Security BearerToken
// Get implements ActivitiesHandler
func (h *activitiesHandler) Get(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.HttpErrorResponse{
			Status:  "Bad Request",
			Message: "Invalid id",
		})
		return
	}

	data, err := h.activitiesService.Get(ctx, id)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, model.HttpErrorResponse{
				Status:  "Not Found",
				Message: fmt.Sprintf("Activity with ID %d Not Found", id),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, model.HttpErrorResponse{
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SingleActivityResponse{
		Status:  "Success",
		Message: "Success",
		Value:   data,
	})
}

// GetListActivities godoc
// @Summary      get all activities
// @Description  get list all of existing activities
// @Tags         activity
// @Produce      json
// @Success      200  {object}  model.ListActivitiesResponse
// @Failure      400  {object}  model.HttpErrorResponse
// @Failure      404  {object}  model.HttpErrorResponse
// @Failure      500  {object}  model.HttpErrorResponse
// @Router       /activity-groups [get]
// @Security BearerToken
// GetAll implements ActivitiesHandler
func (h *activitiesHandler) GetAll(ctx *gin.Context) {
	activities, err := h.activitiesService.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.HttpErrorResponse{
			Status:  "Bad Request",
			Message: err.Error(),
		})
		return
	}

	if activities != nil {
		ctx.JSON(http.StatusOK, model.ListActivitiesResponse{
			Status:  "Success",
			Message: "Success",
			Value:   activities,
		})
		return
	}
	ctx.JSON(http.StatusOK, model.ListActivitiesResponse{
		Status:  "Success",
		Message: "Success",
	})
}

// UpdateActivity godoc
// @Summary      update an activity
// @Description  update an activity for parent of items
// @Tags         activity
// @Accept       json
// @Produce      json
// @Param 		 id   path int true "id of activity"
// @Param		 activity	body		model.CreateActivityParam	true	"Update Activity"
// @Success      200  {object}  model.SingleActivityResponse
// @Failure      400  {object}  model.HttpErrorResponse
// @Failure      404  {object}  model.HttpErrorResponse
// @Failure      500  {object}  model.HttpErrorResponse
// @Router       /activity-groups/{id} [patch]
// @Security BearerToken
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
	err = h.activitiesService.Update(ctx, request)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, model.HttpErrorResponse{
				Status:  "Not Found",
				Message: fmt.Sprintf("Activity with ID %d Not Found", id),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, model.HttpErrorResponse{
			Status:  "Internal Server Error",
			Message: err.Error(),
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
