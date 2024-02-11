package handlers

import (
	"example.com/ideanest-task/pkg/database/mongodb/models"
	"example.com/ideanest-task/pkg/database/mongodb/repository"
	"example.com/ideanest-task/pkg/utils"
	"fmt"
)


func SignUp(user models.User) error{
	user.Password = utils.Hash(user.Password)
	err := repository.CreateUser(user)
	return err
}

func Login(req models.LoginRequest) (string, string, error){
	user,err := repository.GetUser(req.Email)
	if err != nil {
		return "","",err
	}
	if utils.ComparePassword(user.Password,req.Password) {
		access, refresh, err := utils.GenerateToken(req.Email)
		if err != nil {
			return "","",err
		}
		return access, refresh, nil
	}
	return "","",fmt.Errorf("User Not Found")
}

func Refresh(refreshToken string)(string, string, error) {
	access,refresh, err := utils.Refresh(refreshToken)
	if err != nil{
		return "","",nil
	}
	return access,refresh,nil
}