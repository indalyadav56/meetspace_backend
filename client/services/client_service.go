package services

import (
	"fmt"
	"meetspace_backend/client/models"
	"meetspace_backend/client/repositories"
	"meetspace_backend/client/types"
	"meetspace_backend/config"
	"meetspace_backend/user/constants"
	userModel "meetspace_backend/user/models"
	"meetspace_backend/user/service_factory"
	userService "meetspace_backend/user/services"
	"meetspace_backend/utils"

	"github.com/google/uuid"
)


type ClientService struct {
	ClientRepository *repositories.ClientRepository
	UserService  *userService.UserService
}

func NewClientService() *ClientService {
	repo := repositories.NewClientRepository()
	userService := service_factory.GetUserService()
	
	return &ClientService{
		ClientRepository: repo,
		UserService: userService,
	}
}

func (cs *ClientService) CreateClient(clientData types.ClientCreateData) (*models.Client, error) {
	clientObj := models.Client{
		ID: uuid.New(),
		CompanyName: clientData.CompanyName,
		CompanyDomain : clientData.CompanyName+".localhost:3000",
		Country: "india",
	}
	
	client, err := cs.ClientRepository.CreateRecord(clientObj)
	if err != nil {
		fmt.Println("Error creating record", err.Error())
		return nil, err
	}

	hashedPassword, _ := utils.EncryptPassword(clientData.Password)
	userData := userModel.User{
		FirstName: clientData.FirstName,
		LastName: clientData.LastName,
		Email: clientData.Email,
		Password: hashedPassword,
		Role: constants.ROLE_ADMIN,
		IsAdmin: true,
		ClientID: client.ID,
		CreatedBy: client,
		UpdatedBy: client,
	}
	userObj, err := cs.UserService.UserRepository.CreateRecord(userData)
	if err != nil {
		return nil, err
	}
	clientObj.ClientUserID = userObj.ID.String()
	config.DB.Save(&clientObj)

	// send mail to client's given email address
	bodyData := fmt.Sprintf("<h1>Domain</h1>:- <a href=\"http://%s\" target=\"_blank\" rel=\"noopener noreferrer\">%s</a> and password:- %s", clientObj.CompanyDomain, clientObj.CompanyDomain, clientData.Password)

	go utils.SendEmail(
		clientData.Email,
		"sending for password and domain to client",
		bodyData,
	)

	return &clientObj, nil

}

func (cs *ClientService) GetClientById(clientId string) (models.Client, error) {
	client, _ := cs.ClientRepository.GetClientById(clientId)
	return client, nil
}

func (cs *ClientService) GetClientByUserId(userId string) (models.Client, error) {
	client, _ := cs.ClientRepository.GetClientByUserId(userId)
	return client, nil
}

func (cs *ClientService) GetAllClients(companyName string) ([]models.Client, error) {
	clients, err := cs.ClientRepository.GetAllClients(companyName)
	return clients, err
}
