package handlers

import (
	"example.com/ideanest-task/pkg/database/mongodb/models"
	"example.com/ideanest-task/pkg/database/mongodb/repository"
	"fmt"
)

func GetOrganization(id string, email string) (*models.Organization, error){
	orgRepo := repository.NewOrganizationRepository();
	org,err := orgRepo.GetByOrgIdAndUserEmail(id,email)
	if err != nil {
		return nil, err
	}
	return org,err
}

func GetAllOrganizations(email string) ([]models.Organization, error){
	orgRepo := repository.NewOrganizationRepository();
	orgs, err := orgRepo.GetAll(email)
	return orgs,err
}
func CreateOrganization(org models.Organization, email string) (string,error){
	orgRepo := repository.NewOrganizationRepository();
	user, err := repository.GetUser(email)
	if err != nil {
		return "", fmt.Errorf("Unauthorized Access")
	}
	org.OrganizationMembers = append(org.OrganizationMembers, models.Member{Email: user.Email,Name: user.Name,AccessLevel: "Admin"})
	id, err := orgRepo.Create(org)
	if err != nil{
		return "",err
	}
	return id,nil
}

func Update(orgId string,adminEmail string,org models.OrganizationRequest) error{
	orgRepo := repository.NewOrganizationRepository()
	if !orgRepo.IsAdmin(orgId,adminEmail) {
		return fmt.Errorf("UnAuthorized Access")
	}
	err := orgRepo.Update(orgId, org)
	return err
}

func Delete(orgId string, adminEmail string) error {
	orgRepo := repository.NewOrganizationRepository()
	if !orgRepo.IsAdmin(orgId,adminEmail) {
		return fmt.Errorf("UnAuthorized Access")
	}
	err := orgRepo.Delete(orgId)
	return err
}

func AddMember(orgId string, adminEmail string,email string) error {
	user, err := repository.GetUser(email)
	if err != nil {
		return err
	}

	orgRepo := repository.NewOrganizationRepository()
	if !orgRepo.IsAdmin(orgId,adminEmail) {
		return fmt.Errorf("UnAuthorized Access")
	}

	member := models.Member{Email: user.Email,Name: user.Name,AccessLevel: "User"}
	err = orgRepo.AddMember(orgId,member)
	if err != nil {
		return err;
	}
	return nil
}