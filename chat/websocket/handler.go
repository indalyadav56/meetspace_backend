package websocket

import (
	"fmt"
	"log"
	commonServices "meetspace_backend/common/services"
	"meetspace_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(reqWriter http.ResponseWriter, req *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(reqWriter, req, nil)
	
	if err != nil {
		log.Println("Websocket connection error:- ", err)
		return nil, err
	}
	
	return conn, nil
}

type WebSocketHandler struct {
	RedisService *commonServices.RedisService
}

func NewWebSocketHandler(svc *commonServices.RedisService) *WebSocketHandler {
	return &WebSocketHandler{
		RedisService: svc,
	}
}

// gets or create WebSocket pool based on the groupName
func (h *WebSocketHandler) getOrCreatePool(groupName string) *Pool {
	pool, exists := groupPools[groupName]
	if !exists {
		pool = NewPool(h.RedisService)
		groupPools[groupName] = pool
		go pool.Start()
	}
	return pool
}

func (h *WebSocketHandler) handleWebSocketConnection(groupName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentPool := h.getOrCreatePool(groupName)
		h.webSocketServer(currentPool, c)
	}
}

func (h *WebSocketHandler) handleWebSocketConnectionFromParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		groupName := c.Param("groupName")
		currentPool := h.getOrCreatePool(groupName)
		h.webSocketServer(currentPool, c)
	}
}


func (h *WebSocketHandler) webSocketServer(pool *Pool, c *gin.Context)  {
	conn, err := Upgrade(c.Writer, c.Request)
	if err != nil {
		fmt.Println("WebSocket connection error: ", err)
		return
	}
	groupName := c.Param("groupName")
	currentUser, exists := utils.GetUserFromContext(c)

	if !exists{
		return 
	}
	client := &Client{
		Conn:      conn,
		Pool:      pool,
		GroupName: groupName,
		IsGroup:   groupName != "",
		User:      currentUser,
	}
	pool.Register <- client
	client.Read()
}
