package service_factory

import "meetspace_backend/chat/services"


var chatGroupInstance *services.ChatGroupService


func GetChatGroupService() *services.ChatGroupService{
    if chatGroupInstance == nil {
        chatGroupInstance = services.NewChatGroupService()
    }
    
    return chatGroupInstance
}