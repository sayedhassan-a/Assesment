package routes

import (
	"example.com/ideanest-task/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpAuthRoutes(router *gin.Engine){
	router.POST("/signup",controllers.SignUp)
	router.POST("/signin",controllers.Login)
	router.POST("/refresh-token",controllers.RefreshToken)
}
