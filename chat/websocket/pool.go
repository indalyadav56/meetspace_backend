package websocket

import (
	"fmt"
	"meetspace_backend/chat/types"
	"meetspace_backend/utils"
	"sync"
)

var (
	joinedUsers = make(map[string][]string)
	globalClients = make(map[*Client]bool)
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan string
	Service *WebSocketService
	mu sync.Mutex
}

func NewPool() *Pool {
	return &Pool{
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		Clients:     make(map[*Client]bool),
		Broadcast:   make(chan string),
	}
}

func (pool *Pool) Start() {
	for {
		select {
			case client := <-pool.Register:
				pool.registerClient(client)

			case client := <-pool.Unregister:
				pool.unregisterClient(client)

			case payload := <-pool.Broadcast:
				pool.broadcastToClients(payload)
		}
	}
}

func (pool *Pool) registerClient(client *Client) {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	pool.Clients[client] = true

	if client.IsGroup {
		fmt.Println("client.IsGroup: ", client.IsGroup)
		if value, exists := joinedUsers[client.GroupName]; exists {
			value = append(value, client.User.ID.String())
			joinedUsers[client.GroupName] = value
		} else {
			joinedUsers[client.GroupName] = []string{client.User.ID.String()}
		}
	} else {
		globalClients[client] = true
		fmt.Println("globalClients:--->>", globalClients)
		fmt.Println("client.IsGroup:--->>", client.IsGroup)
	}

}

func (pool *Pool) unregisterClient(client *Client) {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	delete(pool.Clients, client)

	if client.IsGroup {
		if users, exists := joinedUsers[client.GroupName]; exists {
			for index, userId := range users {
				if userId == client.User.ID.String() {
					users = append(users[:index], users[index+1:]...)
					joinedUsers[client.GroupName] = users
					break
				}
			}
		}
	}else{
		delete(globalClients, client)
	}
}

func (pool *Pool) broadcastToClients(payload string) {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	for client := range pool.Clients {
		var payloadData types.Payload
		utils.StringToStruct(payload, &payloadData)
		pool.Service.HandleEvent(payloadData, client)

		client.Conn.WriteMessage(1, []byte(payload))
	}
}
