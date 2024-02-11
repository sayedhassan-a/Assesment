package models

type SignUpResponseMessage struct {
	Message			string
}

type LogInResponseMessage struct {
	Message			string	`json:"message" bson:"message"`
	AccessToken		string	`json:"access_token" bson:"access_token"`
	RefreshToken	string	`json:"refresh_token" bson:"refresh_token"`
}

