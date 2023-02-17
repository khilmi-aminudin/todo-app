package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/khilmi-aminudin/todo-app/app/db"
	"github.com/khilmi-aminudin/todo-app/app/router"
	"github.com/khilmi-aminudin/todo-app/handler"
	"github.com/khilmi-aminudin/todo-app/model"
	"github.com/khilmi-aminudin/todo-app/repository"
	"github.com/khilmi-aminudin/todo-app/service"
)

func RunServer() {
	if os.Getenv("GIN_MODE") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	dbconn := db.NewDB()

	dbconn.AutoMigrate(&model.Activities{}, &model.Todos{})

	todosRepo := repository.NewTodosRepository(dbconn)
	todosService := service.NewTodosService(todosRepo)
	todosHandler := handler.NewTodosHandler(todosService)

	activityRepo := repository.NewActivitiesRepository(dbconn)
	activityService := service.NewActivitiesService(activityRepo)
	activityHandler := handler.NewActivitiesHandler(activityService)

	r := gin.Default()
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, model.WebResponse{
			Status:  "success",
			Message: "Hello, I'am Oke",
		})
	})

	router.NewTodosRouter(r, todosHandler)
	router.NewActivitiesRouter(r, activityHandler)

	port := "3030"

	fmt.Println("App is running on port ", port)
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
