package websocket

import (
	"encoding/json"
	"log"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/gorilla/websocket"
)

type Client struct {
	userID uint
	conn   *websocket.Conn
	send   chan models.Message

	logger *log.Logger
}

func NewClient(conn *websocket.Conn, logger *log.Logger, userID uint) *Client {
	return &Client{
		userID: userID,
		conn:   conn,
		send:   make(chan models.Message, 256),
		logger: logger,
	}
}

func (c *Client) readPump(h *Hub) {
	defer func() {
		h.unregister <- c
		c.conn.Close()
	}()

	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket read error (user %d): %v", c.userID, err)
			}
			break
		}
		// Входящие сообщения от клиента не обрабатываются (если не нужно)
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()

	for {
		message, ok := <-c.send
		if !ok {
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		w, err := c.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Printf("Failed to get writer (user %d): %v", c.userID, err)
			return
		}

		if err := json.NewEncoder(w).Encode(message); err != nil {
			log.Printf("Failed to encode message (user %d): %v", c.userID, err)
			return
		}

		if err := w.Close(); err != nil {
			log.Printf("Failed to close writer (user %d): %v", c.userID, err)
			return
		}
	}
}
