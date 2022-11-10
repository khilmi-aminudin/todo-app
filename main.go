package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khilmi-aminudin/todo-app/db"
	"github.com/khilmi-aminudin/todo-app/handler"
	"github.com/khilmi-aminudin/todo-app/repository"
	"github.com/khilmi-aminudin/todo-app/service"
)

func main() {
	var (
		db                 *sql.DB                       = db.NewDB()
		activityRepository repository.ActivityRepository = repository.NewActivityRepository(db)
		activityService    service.ActivityService       = service.NewActivityService(activityRepository)
		activityHandler    handler.ActivityHandler       = handler.NewActivityHandler(activityService)

		todoRepository repository.TodoRepository = repository.NewTodoRepository(db)
		todoService    service.TodoService       = service.NewTodoService(todoRepository)
		todoHandler    handler.TodoHandler       = handler.NewTodoHandler(todoService)
	)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/api/activity-group/:id", activityHandler.Get)
	r.GET("/api/activity-group", activityHandler.GetAll)
	r.POST("/api/activity-group", activityHandler.Create)
	r.DELETE("/api/activity-group/:id", activityHandler.Delete)
	r.PATCH("/api/activity-group/:id", activityHandler.Update)

	r.POST("/api/todo-item", todoHandler.Create)

	r.Run(":8080")
}
