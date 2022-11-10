package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khilmi-aminudin/todo-app/service"
	"github.com/khilmi-aminudin/todo-app/web"
)

type ActivityHandler interface {
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type activityHandler struct {
	service service.ActivityService
}

func NewActivityHandler(service service.ActivityService) ActivityHandler {
	return &activityHandler{
		service: service,
	}
}

func (handler *activityHandler) Get(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)

	if err != nil {
		panic(err.Error())
	}

	response := handler.service.Get(int64(id))

	c.JSON(http.StatusOK, response)

}

func (handler *activityHandler) GetAll(c *gin.Context) {
	response := handler.service.GetAll()

	c.JSON(http.StatusOK, response)
}

func (handler *activityHandler) Create(c *gin.Context) {
	var requestBody web.ActivityCreateRequest

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		fmt.Printf("error : %v\n", err)

		if err.Error() == "Key: 'ActivityCreateRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag" {
			c.JSON(http.StatusBadRequest, web.WebResponse{
				Status:  "Bad Request",
				Message: "email cannt be null",
			})
			return
		}

		if err.Error() == "Key: 'ActivityCreateRequest.Title' Error:Field validation for 'Title' failed on the 'required' tag" {
			c.JSON(http.StatusBadRequest, web.WebResponse{
				Status:  "Bad Request",
				Message: "title cannt be null",
			})
			return
		}

		if err.Error() == "Key: 'ActivityCreateRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag" {
			c.JSON(http.StatusBadRequest, web.WebResponse{
				Status:  "Bad Request",
				Message: "invalid email format",
			})
			return
		}

		if err.Error() == "Key: 'ActivityCreateRequest.Title' Error:Field validation for 'Title' failed on the 'min' tag" {
			c.JSON(http.StatusBadRequest, web.WebResponse{
				Status:  "Bad Request",
				Message: "title min 5 character",
			})
			return
		}
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Status:  "Bad Request",
			Message: "invalid email and title",
		})
		return
	}

	response := handler.service.Create(requestBody)

	c.JSON(http.StatusCreated, response)

}

func (handler *activityHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Status:  "Bad Request",
			Message: "Invalid ID",
		})
	}

	var reqBody web.ActivityUpdateRequest

	err = c.ShouldBindJSON(&reqBody)

	reqBody.ID = int64(id)

	if err != nil {
		panic(err.Error())
	}

	response := handler.service.Update(reqBody)

	if response.Status == "Not Found" {
		c.JSON(http.StatusNotFound, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (handler *activityHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.WebResponse{
			Status:  "Bad Request",
			Message: "Invalid ID",
		})
	}
	response := handler.service.Delete(int64(id))

	c.JSON(http.StatusOK, response)
}
