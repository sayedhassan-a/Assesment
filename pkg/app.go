package pkg

import (
	"example.com/ideanest-task/pkg/api/routes"
	"example.com/ideanest-task/pkg/utils"
	"github.com/gin-gonic/gin"
)


func Run()  {
	utils.ConnectToDB()
	router := gin.Default()
	routes.SetUpAuthRoutes(router)
	routes.SetUpOrgRoutes(router)
	router.Run("0.0.0.0:8080")
}