package service

import (
	"fmt"
	"time"

	"github.com/khilmi-aminudin/todo-app/models"
	"github.com/khilmi-aminudin/todo-app/repository"
	"github.com/khilmi-aminudin/todo-app/web"
)

type ActivityService interface {
	Get(id int64) web.WebResponse
	GetAll() web.WebResponse
	Create(activity web.ActivityCreateRequest) web.WebResponse
	Update(activity web.ActivityUpdateRequest) web.WebResponse
	Delete(id int64) web.WebResponse
}

type activityService struct {
	repository repository.ActivityRepository
}

func NewActivityService(repo repository.ActivityRepository) ActivityService {
	return &activityService{
		repository: repo,
	}
}

func (service *activityService) Get(id int64) web.WebResponse {
	data := service.repository.Get(models.Activity{ID: id})

	var empty models.Activity

	var response web.WebResponse

	if data == empty {
		response.Status = "Not Found"
		response.Message = fmt.Sprintf("activity with id %d not found", id)
	} else {
		response.Status = "Success"
		response.Message = "Success"
		response.Data = data
	}

	return response

}

func (service *activityService) GetAll() web.WebResponse {

	data := service.repository.GetAll()

	response := web.WebResponse{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}

	return response
}

func (service *activityService) Create(activity web.ActivityCreateRequest) web.WebResponse {

	params := models.Activity{
		Email:     activity.Email,
		Title:     activity.Title,
		CreatedAt: time.Now(),
	}

	data := service.repository.Create(params)

	return web.WebResponse{
		Status:  "success",
		Message: "Activity Created",
		Data:    data,
	}
}

func (service *activityService) Update(activity web.ActivityUpdateRequest) web.WebResponse {

	dataExist := service.repository.Get(models.Activity{ID: activity.ID})

	var empty models.Activity

	if dataExist == empty {
		return web.WebResponse{
			Status:  "Not Found",
			Message: "Activity Not Found",
		}
	}

	updateParams := models.Activity{ID: activity.ID, Email: activity.Email, Title: activity.Title, UpdatedAt: time.Now()}
	data := service.repository.Update(updateParams)

	return web.WebResponse{
		Status:  "Success",
		Message: "Activity Updated",
		Data:    data,
	}

}

func (service *activityService) Delete(id int64) web.WebResponse {

	activity := service.repository.Get(models.Activity{ID: id})

	var response web.WebResponse

	empty := models.Activity{}
	if activity == empty {
		response.Status = "Not Found"
		response.Message = fmt.Sprintf("Activity with id %d not exist", id)
	}
	err := service.repository.Delete(models.Activity{ID: id})

	if err != nil {
		response.Status = "Error"
		response.Message = "Internal Server Error"
	}

	response.Status = "Success"
	response.Message = fmt.Sprintf("Activity with id %d deleted", id)
	return response
}
