package router

import (
	"github.com/gin-gonic/gin"

	"github.com/khilmi-aminudin/todo-app/handler"
)

func NewActivitiesRouter(r *gin.RouterGroup, activitieshandler handler.ActivitiesHandler) {
	r.POST("/activity-groups", activitieshandler.Create)
	r.PATCH("/activity-groups/:id", activitieshandler.Update)
	r.DELETE("/activity-groups/:id", activitieshandler.Delete)
	r.GET("/activity-groups", activitieshandler.GetAll)
	r.GET("/activity-groups/:id", activitieshandler.Get)
}
