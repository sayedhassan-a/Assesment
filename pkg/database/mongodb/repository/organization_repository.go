package repository

import (
	"context"
	"example.com/ideanest-task/pkg/database/mongodb/models"
	"example.com/ideanest-task/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type OrganizationRepository struct {
	collection	*mongo.Collection
}

func NewOrganizationRepository()*OrganizationRepository  {
	collection := utils.GetCollection(utils.DB,"orgs")
	return &OrganizationRepository{
		collection: collection,
	}
}

func (repository *OrganizationRepository) GetAll(email string) ([]models.Organization,error) {
	cursor, err := repository.collection.Find(context.Background(),bson.M{"organization_members":bson.M{"$elemMatch":bson.M{"email":email}}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var orgs []models.Organization
	if err := cursor.All(context.Background(), &orgs); err != nil {
		return nil, err
	}
	return orgs,nil
}

func (repository *OrganizationRepository) Create(org models.Organization) (string,error) {
	ctx, cancel	 := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	result,err := repository.collection.InsertOne(ctx,org)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (repository *OrganizationRepository) GetByOrgIdAndUserEmail(orgId string, email string) (*models.Organization,error) {
	ctx, cancel	 := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(orgId)
	if err != nil {
		return nil, err
	}
	var org models.Organization
	err = repository.collection.FindOne(ctx,bson.M{"_id":id,"organization_members":bson.M{"$elemMatch": bson.M{"email":email}}}).Decode(&org)
	if err != nil {
		return nil, err
	}
	return &org, err
}

func (repository *OrganizationRepository) Update(orgId string, org models.OrganizationRequest) error {
	ctx, cancel	 := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(orgId)
	if err != nil {
		return err
	}
	_,err = repository.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set":org})
	return err
}

func (repository *OrganizationRepository) AddMember(orgId string, member models.Member) error {
	ctx, cancel	 := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(orgId)
	if err != nil {
		return err
	}
	_,err = repository.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$push": bson.M{"organization_members": member}})
	return err
}

func (repository *OrganizationRepository) Delete(orgId string) error {
	ctx, cancel	 := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	id, err := primitive.ObjectIDFromHex(orgId)
	if err != nil {
		return err
	}
	_,err = repository.collection.DeleteOne(ctx,bson.M{"_id":id})
	if err != nil {
		return err
	}
	return err
}

func (repository *OrganizationRepository) IsAdmin(orgId string, email string) bool {
	ctx, cancel	 := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	var org models.Organization
	id, err := primitive.ObjectIDFromHex(orgId)
	if err != nil {
		return false
	}
	err = repository.collection.FindOne(ctx,bson.M{"_id":id,"organization_members":bson.M{"$elemMatch": bson.M{"email":email,"access_level":"Admin"}}}).Decode(&org)
	if err != nil {
		return false
	}
	return true
}