package websocket

import (
	"meetspace_backend/chat/constants"
	"meetspace_backend/chat/models"
	"meetspace_backend/chat/types"
	"meetspace_backend/config"
	"meetspace_backend/utils"
	// userService "meetspace_backend/internal/user/services"
)

type WebSocketService struct {}

func NewWebSocketService() *WebSocketService {
	return &WebSocketService{}
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
	// currentRoom, err := config.ChatRoomService.GetChatRoomByID(client.GroupName)
	
	// if chat room not found then create a new chat room for sender and receiver user
	// if err != nil {
	// 	receiverUserData := payload.Data["receiver_user"].(map[string]interface{})
	// 	var users []string
	// 	users = append(users, receiverUserData["id"].(string))
	// 	// config.ChatRoomService.CreateChatRoomRecord("NewChatRoom", client.User.ID.String(), users)
	// 	CheckMessageNotification(client, payload)
	// }else{
	// 	// senderUserData := payload.Data["sender"].(map[string]interface{})
	// 	// config.ChatMessageService.CreateChatMessage("NewChatMessageContent", senderUserData["id"].(string), currentRoom.ID.String())
	// 	CheckMessageNotification(client, payload)
	// }
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
			// send notification
			notificationData:= map[string]interface{}{
				"receiver_user": map[string]interface{}{
					"id": roomUserId,
				},
				"sender_user": payload.Data["sender"],
				"room_id": chatRoomObj.ID.String(),
				"is_group": chatRoomObj.IsGroup,
				"content": payload.Data["content"],
				"room_name": chatRoomObj.RoomName,
			}
			stringData := SendChatMessageNotification(notificationData)
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