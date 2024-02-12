package websocket

import (
	"meetspace_backend/chat/constants"
	"meetspace_backend/chat/models"
	"meetspace_backend/chat/services"
	"meetspace_backend/chat/types"
	commonServices "meetspace_backend/common/services"
	"meetspace_backend/config"
	userServices "meetspace_backend/user/services"
	"meetspace_backend/utils"
)

type WebSocketService struct {
	LoggerService *commonServices.LoggerService
	ChatRoomService *services.ChatRoomService
	ChatMessageService *services.ChatMessageService
	UserService *userServices.UserService
}

func NewWebSocketService(loggerService *commonServices.LoggerService, chatRoomSvc *services.ChatRoomService, chatMsgSvc *services.ChatMessageService, userSvc *userServices.UserService) *WebSocketService {
	return &WebSocketService{
		LoggerService: loggerService,
		ChatRoomService: chatRoomSvc,
		ChatMessageService: chatMsgSvc,
		UserService: userSvc,
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
	mapData, _ := utils.StructToMap(payload.Data)
	currentRoom, err := ws.ChatRoomService.GetChatRoomByID(client.GroupName)
	// if chat room not found then create a new chat room for sender and receiver user
	if err != nil {
		var users []string
		receiverUser := mapData["receiver_user"].(map[string]interface{})
		users = append(users, receiverUser["id"].(string))
		room, _ := ws.ChatRoomService.CreateChatRoom("NewChatRoom", client.User.ID.String(), users)
		ws.ChatMessageService.CreateChatMessage(mapData["content"].(string),  client.User.ID.String(), room.ID.String())
		
	}else{
		ws.ChatMessageService.CreateChatMessage(mapData["content"].(string),  client.User.ID.String(), currentRoom.ID.String())
	}

	CheckMessageNotification(client, payload)
}

func (ws *WebSocketService) HandleChatNotificationReceived(payload types.Payload, client *Client) {
	// Implement logic for chat notification received event
}

func CheckMessageNotification(client *Client, payload types.Payload){
	users, exists := joinedUsers[client.GroupName]

	if !exists{
		return
	}
	
	var chatRoomObj models.ChatRoom

	config.DB.Preload("RoomUsers").Where("id=?", client.GroupName).Find(&chatRoomObj)

	joinedUsersMap := make(map[string]bool)
	for _, userId := range users {
		joinedUsersMap[userId] = true
	}

	for _, userObj := range chatRoomObj.RoomUsers {
		roomUserId := userObj.ID.String()
		
		if !joinedUsersMap[roomUserId] {
			mapData, _ := utils.StructToMap(payload.Data)
			stringData := SendChatMessageNotification(mapData)
			SendMessageToClients(globalClients, stringData)

		}
	}

	return
}


func SendChatMessageNotification(notificationData map[string]interface{}) string{
	newEventData := types.Payload{
		Event: constants.CHAT_NOTIFICATION_SENT,
		Data: notificationData,
	}
	data, _ := utils.StructToString(newEventData)
	return data
}


// broadcastMessage to all the given clients
func SendMessageToClients(clients map[*Client]bool, msgData string){
	for client := range clients{
		client.Conn.WriteMessage(1, []byte(msgData))
	}
}