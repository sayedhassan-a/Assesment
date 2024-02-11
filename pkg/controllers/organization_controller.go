package controllers

import (
	"example.com/ideanest-task/pkg/api/handlers"
	"example.com/ideanest-task/pkg/database/mongodb/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrganization(ctx *gin.Context) {
	var org models.Organization
	if err := ctx.ShouldBindJSON(&org); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := handlers.CreateOrganization(org,ctx.GetString("email"))
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"organization_id":id})
}

func GetAllOrganizations(ctx *gin.Context) {
	orgs, err := handlers.GetAllOrganizations(ctx.GetString("email"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,orgs)
}
func GetOrganizationById(ctx *gin.Context) {
	id := ctx.Param("id")
	email := ctx.GetString("email")
	res, err := handlers.GetOrganization(id,email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,*res)

}

func UpdateOrganization(ctx *gin.Context) {
	var orgRequest models.OrganizationRequest
	if err := ctx.ShouldBindJSON(&orgRequest); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := handlers.Update(ctx.Param("id"),ctx.GetString("email"),orgRequest)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"organization_id":ctx.Param("id"),"name":orgRequest.Name,"description":orgRequest.Description})
}

func DeleteOrganization(ctx *gin.Context) {
	err := handlers.Delete(ctx.Param("id"),ctx.GetString("email"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message":"success"})
}


func AddMember(ctx *gin.Context) {
	var requestBody struct
	{
		Email string
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id := ctx.Param("id")
	adminEmail := ctx.GetString("email")
	err := handlers.AddMember(id,adminEmail,requestBody.Email)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message":"success"})
}