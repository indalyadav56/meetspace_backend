package service_factory

import "meetspace_backend/user/services"


var userServiceInstance *services.UserService


func GetUserService() *services.UserService{
	
    if userServiceInstance == nil {
        userServiceInstance = services.NewUserService()
    }
    
    return userServiceInstance
}