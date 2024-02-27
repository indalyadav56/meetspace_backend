package websocket

import (
	"meetspace_backend/chat/constants"
	"meetspace_backend/chat/services"
	"meetspace_backend/chat/types"
	commonServices "meetspace_backend/common/services"
	userServices "meetspace_backend/user/services"
)

type WebSocketService struct {
	LoggerService *commonServices.LoggerService
	ChatRoomService *services.ChatRoomService
	ChatMessageService *services.ChatMessageService
	UserService *userServices.UserService
	RedisService *commonServices.RedisService
}

func NewWebSocketService(loggerService *commonServices.LoggerService, chatRoomSvc *services.ChatRoomService, chatMsgSvc *services.ChatMessageService, userSvc *userServices.UserService, redisService *commonServices.RedisService) *WebSocketService {
	return &WebSocketService{
		LoggerService: loggerService,
		ChatRoomService: chatRoomSvc,
		ChatMessageService: chatMsgSvc,
		UserService: userSvc,
		RedisService: redisService,
	}
}

func (ws *WebSocketService) HandleEvent(payload types.Payload, client *Client) {
	switch payload.Event {
		case constants.USER_CONNECTED:
			ws.HandleUserConnected(payload, client)

		case constants.USER_DISCONNECTED:
			ws.HandleUserDisconnected(payload, client)

		case constants.CHAT_MESSAGE_SENT:
			ws.HandleChatMessageSent(payload, client)

		case constants.CHAT_NOTIFICATION_RECEIVED:
			ws.HandleChatNotificationReceived(payload, client)

		default:
			// Handle other events or log unsupported events
	}
}

func (ws *WebSocketService) HandleUserConnected(payload types.Payload, client *Client) {
	// Implement logic for user connected event
}

func (ws *WebSocketService) HandleUserDisconnected(payload types.Payload, client *Client) {
	// Implement logic for user disconnected event
}

func (ws *WebSocketService) HandleChatMessageSent(payload types.Payload, client *Client) {
	// CheckMessageNotification(client, payload)
}

func (ws *WebSocketService) HandleChatNotificationReceived(payload types.Payload, client *Client) {
	// Implement logic for chat notification received event
}
