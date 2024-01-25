package service_factory

import "meetspace_backend/client/services"


var clientServiceInstance *services.ClientService
var clientUserServiceInstance *services.ClientUserService


func GetClientService() *services.ClientService{
	if clientServiceInstance == nil {
        clientServiceInstance = services.NewClientService()
    }
    return clientServiceInstance
}

func GetClientUserService() *services.ClientUserService{
	if clientUserServiceInstance == nil {
        clientUserServiceInstance = services.NewClientUserService()
    }
    return clientUserServiceInstance
}