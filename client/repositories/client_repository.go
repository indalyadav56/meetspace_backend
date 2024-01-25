package repositories

import (
	"meetspace_backend/client/models"
	"meetspace_backend/config"

	"gorm.io/gorm"
)

type IClientRepository interface {
    CreateRecord(clientModel models.Client) (models.Client, error)
    GetClientById(clientId string) (models.Client, error) 
}


type ClientRepository struct{
	DB *gorm.DB
}


func NewClientRepository() *ClientRepository {
	return &ClientRepository{
		DB: config.DB,
	}
}


func (c *ClientRepository) CreateRecord(clientModel models.Client) (*models.Client, error){
	result := c.DB.Create(&clientModel)
	if result.Error != nil {
        return nil, result.Error
    }
	return &clientModel, nil
}


func (c *ClientRepository) GetClientById(clientId string) (models.Client, error){
	var client models.Client
	result := c.DB.Where("id = ?", clientId).First(&client)
	if result.Error != nil {
        return client, result.Error
    }
	return client, nil
}

func (c *ClientRepository) GetClientByUserId(userId string) (models.Client, error){
	var client models.Client
	c.DB.Where("client_user_id = ?", userId).First(&client)
	return client, nil
}

func (c *ClientRepository) GetAllClients(companyName string) ([]models.Client, error){
	var clients []models.Client
	c.DB.Where("company_name = ?", companyName).Find(&clients)
	return clients, nil
}