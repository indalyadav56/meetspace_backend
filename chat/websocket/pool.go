package websocket

import (
	"context"
	"fmt"
	"meetspace_backend/chat/constants"
	"meetspace_backend/chat/types"
	"meetspace_backend/utils"
	"sync"

	"github.com/go-redis/redis/v8"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan string
	Service *WebSocketService
	mu sync.Mutex
}

func NewPool(svc *WebSocketService) *Pool {
	return &Pool{
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		Clients:     make(map[*Client]bool),
		Broadcast:   make(chan string),
		Service: svc,
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

	if client.GroupName != ""{
		currentGroup := fmt.Sprintf("client:group:%v", client.GroupName)
	
		pool.Service.RedisService.SAdd(currentGroup, client.User.ID.String())
		fmt.Println("this user added in client:group, redis sets")

		pubsub := pool.Service.RedisService.Subscribe(currentGroup)
		handleRedisMessages(pubsub, pool)
		
	}else{
		pool.Service.RedisService.SAdd("user:online", client.User.ID.String())

		pubsub := pool.Service.RedisService.Subscribe("client")
		handleRedisMessages(pubsub, pool)

		// publish connected user to clients
		payload := types.Payload{
			Event: constants.USER_CONNECTED,
			Data: map[string]interface{}{
				"id": client.User.ID.String(),
			},
		}
		strData, _ := utils.StructToString(payload)
		pool.Service.RedisService.Publish("client", strData)
	}
}

func (pool *Pool) unregisterClient(client *Client) {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	
	if client.GroupName != ""{
		currentGroup := fmt.Sprintf("client:group:%v", client.GroupName)
		fmt.Println("groupname", currentGroup)
		pool.Service.RedisService.SRem(currentGroup, client.User.ID.String())
		fmt.Println("client:group deleted from redis sets")
	}else{
		// publish disconnect user to clients
		payload := types.Payload{
			Event: constants.USER_DISCONNECTED,
			Data: map[string]interface{}{
				"id": client.User.ID.String(),
			},
		}
		strData, _ := utils.StructToString(payload)
		pool.Service.RedisService.Publish("client", strData)

		// remove disconnected user
		pool.Service.RedisService.SRem("user:online", client.User.ID.String())
	}
	
	fmt.Println("client unregister successfully")
}

func (pool *Pool) broadcastToClients(payload string) {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	for client := range pool.Clients {
		var payloadData types.Payload
		utils.StringToStruct(payload, &payloadData)
		client.Conn.WriteMessage(1, []byte(payload))
		// pool.Service.HandleEvent(payloadData, client)
	}
}

func handleRedisMessages(pubsub *redis.PubSub, pool *Pool) {
    go func() {
        // Continuously receive messages from Redis Pub/Sub
        for {
            msg, err := pubsub.ReceiveMessage(context.Background())
            if err != nil {
                panic(err) // Handle errors appropriately in production
            }

            // Forward received messages to all connected websocket clients
            payload := msg.Payload
            pool.Broadcast <- payload
        }
    }()
}
