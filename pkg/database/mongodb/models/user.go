package models

type User struct {
	Name		string	`json:"name" bson:"name"`
	Email		string	`json:"email" validate:"email,required" bson:"email"`
	Password	string	`json:"password" validate:"required,min=6" bson:"password"`
}
