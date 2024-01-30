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
	UserService = userServices.NewUserService(userRepo.NewUserRepository(GetDB()))
	AuthService = authServices.NewAuthService(userServices.NewUserService(userRepo.NewUserRepository(GetDB())))
	VerificationService = authServices.NewVerificationService(authRepo.NewVerificationRepository(GetDB()))
	ClientService = clientServices.NewClientService(clientRepo.NewClientRepository(GetDB()), UserService)
	ClientUserService = clientServices.NewClientUserService(clientRepo.NewClientRepository(GetDB()), UserService)
	ChatRoomService = chatServices.NewChatRoomService(chatRepo.NewChatRoomRepository(GetDB()), UserService)
	ChatGroupService = chatServices.NewChatGroupService(chatRepo.NewChatRoomRepository(GetDB()), UserService)
	ChatMessageService = chatServices.NewChatMessageService(chatRepo.NewChatMessageRepository(GetDB()), UserService, ChatRoomService)
}