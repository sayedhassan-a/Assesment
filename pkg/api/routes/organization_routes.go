package routes

import (
	"example.com/ideanest-task/pkg/api/middleware"
	"example.com/ideanest-task/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpOrgRoutes(router *gin.Engine){
	org := router.Group("/organization")
	org.Use(middleware.Authenticate)
	org.GET(":id",controllers.GetOrganizationById)
	org.POST("/",controllers.CreateOrganization)
	org.PUT("/:id",controllers.UpdateOrganization)
	org.GET("/",controllers.GetAllOrganizations)
	org.DELETE("/:id",controllers.DeleteOrganization)
	org.POST("/:id/invite",controllers.AddMember)

}
