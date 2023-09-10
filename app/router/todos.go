package router

import (
	"github.com/gin-gonic/gin"

	"github.com/khilmi-aminudin/todo-app/handler"
)

func NewTodosRouter(r *gin.RouterGroup, todoshandler handler.TodosHandler) {
	r.POST("/todo-items", todoshandler.Create)
	r.PATCH("/todo-items/:id", todoshandler.Update)
	r.DELETE("/todo-items/:id", todoshandler.Delete)
	r.GET("/todo-items", todoshandler.GetAll)
	r.GET("/todo-items/:id", todoshandler.Get)
}
