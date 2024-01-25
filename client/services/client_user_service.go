package services

import (
	"meetspace_backend/client/repositories"
	"meetspace_backend/client/types"
	userModel "meetspace_backend/user/models"
	"meetspace_backend/user/service_factory"
	userService "meetspace_backend/user/services"
	"meetspace_backend/utils"
)


type ClientUserService struct {
	ClientRepository *repositories.ClientRepository
	UserService  *userService.UserService
}

func NewClientUserService() *ClientUserService {
	repo := repositories.NewClientRepository()
	userService := service_factory.GetUserService()
	
	return &ClientUserService{
		ClientRepository: repo,
		UserService: userService,
	}
}

func (cs *ClientUserService) AddClientUser(clientData types.ClientAddUser) (*userModel.User, error) {
	hashedPass, _ := utils.EncryptPassword(clientData.Password)
	userData := userModel.User{
		FirstName: clientData.FirstName,
		LastName: clientData.LastName,
		Email: clientData.Email,
		Password: hashedPass,
		ClientID: clientData.ClientID,
		CreatedBy: clientData.CreatedBy,
		UpdatedBy: clientData.UpdatedBy,
	}
	
	userObj, err := cs.UserService.UserRepository.CreateRecord(userData)
	if err != nil {
		return nil, err
	}

	return userObj, nil
}

func (cs *ClientUserService) GetClientUsers(clientId string) ([]userModel.User, error) {
	users, _ := cs.UserService.UserRepository.GetUsersByClientId(clientId)
	return users, nil
}
