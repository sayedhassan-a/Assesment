package controllers

import (
	"example.com/ideanest-task/pkg/api/handlers"
	"example.com/ideanest-task/pkg/database/mongodb/models"
	"example.com/ideanest-task/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)



func SignUp(ctx *gin.Context){

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := handlers.SignUp(user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func Login(ctx *gin.Context){

	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	accessToken,refreshToken, err := handlers.Login(req)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	res := models.LogInResponseMessage{Message: "success",AccessToken: accessToken,RefreshToken: refreshToken}
	ctx.JSON(http.StatusOK, res)

}

func RefreshToken(ctx *gin.Context){

	var refreshToken models.RefreshToken
	if err := ctx.ShouldBindJSON(&refreshToken); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	access, refresh, err := utils.Refresh(refreshToken.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.LogInResponseMessage{
		Message: "sucecss",
		AccessToken: access,
		RefreshToken: refresh,
	})




}