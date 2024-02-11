package repository

import (
	"context"
	"example.com/ideanest-task/pkg/database/mongodb/models"
	"example.com/ideanest-task/pkg/utils"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user models.User) error {
	collection := utils.GetCollection(utils.DB,"users")
	_,err := GetUser(user.Email)
	if err != nil {
		_,err := collection.InsertOne(context.Background(),user)
		return err
	}
	return fmt.Errorf("There is an account with this email")
}

func GetUser(email string) (models.User,error) {
	collection := utils.GetCollection(utils.DB,"users")
	var user models.User
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
