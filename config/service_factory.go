package config

import (
	authRepo "meetspace_backend/auth/repositories"
	authServices "meetspace_backend/auth/services"
	chatRepo "meetspace_backend/chat/repositories"
	chatServices "meetspace_backend/chat/services"
	clientRepo "meetspace_backend/client/repositories"
	clientServices "meetspace_backend/client/services"
	userRepo "meetspace_backend/user/repositories"
	userServices "meetspace_backend/user/services"
)

var (
	AuthService *authServices.AuthService
	VerificationService *authServices.VerificationService
	UserService *userServices.UserService
	ClientService *clientServices.ClientService
	ClientUserService *clientServices.ClientUserService
	ChatRoomService *chatServices.ChatRoomService
	ChatGroupService *chatServices.ChatGroupService
	ChatMessageService *chatServices.ChatMessageService
)

func StartService(){
	db := GetDB()
	UserService = userServices.NewUserService(userRepo.NewUserRepository(db))
	VerificationService = authServices.NewVerificationService(authRepo.NewVerificationRepository(db))
	ClientService = clientServices.NewClientService(clientRepo.NewClientRepository(db), UserService)
	ClientUserService = clientServices.NewClientUserService(clientRepo.NewClientRepository(db), UserService)
	ChatRoomService = chatServices.NewChatRoomService(chatRepo.NewChatRoomRepository(db), UserService)
	ChatGroupService = chatServices.NewChatGroupService(chatRepo.NewChatRoomRepository(db), UserService)
	ChatMessageService = chatServices.NewChatMessageService(chatRepo.NewChatMessageRepository(db), UserService, ChatRoomService)
}