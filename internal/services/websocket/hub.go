package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/gorilla/websocket"
)

type Hub struct {
	clients    map[uint]*Client
	broadcast  chan models.Message
	register   chan *Client
	unregister chan *Client
	mutex      sync.Mutex

	logger *log.Logger
}

func NewHub(logger *log.Logger) *Hub {
	return &Hub{
		broadcast:  make(chan models.Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[uint]*Client),

		logger: logger,
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.mutex.Lock()
			h.clients[client.userID] = client
			h.mutex.Unlock()
		case client := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.clients[client.userID]; ok {
				delete(h.clients, client.userID)
				close(client.send)
			}
			h.mutex.Unlock()
		case message := <-h.broadcast:
			h.mutex.Lock()
			for _, client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client.userID)
				}
			}
			h.mutex.Unlock()
		}
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // В продакшене следует проверять origin!
	},
}

func (h *Hub) ServeRegister(w http.ResponseWriter, r *http.Request, userId uint) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Printf("cant upgrade request to ws: %s", err)
		return err
	}
	h.logger.Printf("upgrade to websocket")

	client := NewClient(conn, h.logger, userId)

	h.register <- client

	go client.writePump()
	go client.readPump(h)
	// go client.testMessages()

	h.logger.Printf("added new subscriber: %d", userId)

	return nil
}

func (h *Hub) ServeUnregister(w http.ResponseWriter, r *http.Request, userId uint) {
	if _, ok := h.clients[userId]; ok {
		h.unregister <- h.clients[userId]
	}
}

func (h *Hub) AddBroadcastMessage(message models.Message) {
	h.broadcast <- message
}

func InvokeHub(hub *Hub) {
	go hub.run()
}
