package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Member struct {
	Name			string	`json:"name" bson:"name"`
	Email			string	`json:"email" bson:"email"`
	AccessLevel		string	`json:"access_level" bson:"access_level"`
}
type Organization struct {
	Id						primitive.ObjectID	`json:"organization_id" bson:"_id,omitempty"`
	Name					string				`json:"name" bson:"name"`
	Description				string				`json:"description" bson:"description"`
	OrganizationMembers		[]Member 			`json:"organization_members" bson:"organization_members"`
}
