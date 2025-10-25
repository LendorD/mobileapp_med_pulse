package websocket

import (
	"encoding/json"
	"log"
	"sync/atomic"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/gorilla/websocket"
)

type Client struct {
	userID   uint
	conn     *websocket.Conn
	send     chan models.Message
	lastPong int64
	logger   *log.Logger
}

func NewClient(conn *websocket.Conn, logger *log.Logger, userID uint) *Client {
	return &Client{
		userID:   userID,
		conn:     conn,
		send:     make(chan models.Message, 256),
		lastPong: time.Now().UnixNano(),
		logger:   logger,
	}
}

func (c *Client) setLastPong() {
	atomic.StoreInt64(&c.lastPong, time.Now().UnixNano())
}

// Возвращает время последнего pong
func (c *Client) getLastPong() time.Time {
	last := atomic.LoadInt64(&c.lastPong)
	return time.Unix(0, last)
}

func (c *Client) readPump(h *Hub) {
	defer func() {
		h.unregister <- c
		c.conn.Close()
	}()

	// Устанавливаем обработчик Pong-сообщений
	c.conn.SetPongHandler(func(appData string) error {
		c.setLastPong()
		c.logger.Printf("Pong received from user %d", c.userID)
		return nil
	})

	for {
		messageType, _, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.logger.Printf("WebSocket read error (user %d): %v", c.userID, err)
			}
			break
		}

		// Игнорируем входящие текстовые сообщения
		// Ping/Pong обрабатываются автоматически через SetPongHandler
		_ = messageType
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				c.logger.Printf("Failed to get writer (user %d): %v", c.userID, err)
				return
			}
			if err := json.NewEncoder(w).Encode(message); err != nil {
				c.logger.Printf("Failed to encode message (user %d): %v", c.userID, err)
				return
			}
			if err := w.Close(); err != nil {
				c.logger.Printf("Failed to close writer (user %d): %v", c.userID, err)
				return
			}

		case <-ticker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.logger.Printf("Failed to send ping (user %d): %v", c.userID, err)
				return
			}
		}
	}
}
