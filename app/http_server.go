package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

	dbGorm := db.NewGormDB()
	err := dbGorm.AutoMigrate(&model.Activities{}, &model.Todos{})
	if err != nil {
		log.Fatal("error migrating models : ", err)
	}

	dbConn := db.NewDB()

	todosRepo := repository.NewTodosRepository(dbConn)
	todosService := service.NewTodosService(todosRepo)
	todosHandler := handler.NewTodosHandler(todosService)

	activityRepo := repository.NewActivitiesRepository(dbConn)
	activityService := service.NewActivitiesService(activityRepo)
	activityHandler := handler.NewActivitiesHandler(activityService)

	r := gin.Default()
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, model.WebResponse{
			Status:  "success",
			Message: "Hello, I'am Oke",
		})
	})

	baseRoute := r.Group("/api/v1")

	router.NewTodosRouter(baseRoute, todosHandler)
	router.NewActivitiesRouter(baseRoute, activityHandler)

	port := os.Getenv("APP_PORT")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Println("App is running on port ", port)
	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
