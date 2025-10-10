package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	logging "gitlab.com/devkit3/logger"
	"gitlab.com/enterprisemes/notification-service/internal/domain/models"
	"gitlab.com/enterprisemes/notification-service/internal/service/services"
)

type Client struct {
	userID uint
	conn   *websocket.Conn
	send   chan models.Message

	groupIDs []uint

	logger *logging.Logger
	auth   *services.AuthService
}

func NewClient(conn *websocket.Conn, logger *logging.Logger, userId uint, auth *services.AuthService) (*Client, error) {
	groups, _, err := auth.API.GetUserGroups(auth.Token, int(userId))
	if err != nil {
		return nil, fmt.Errorf("cant get groups: %s", err)
	}

	groupsIDs := make([]uint, 0)
	for _, group := range groups {
		groupsIDs = append(groupsIDs, uint(group.Id))
	}

	return &Client{
		userID: userId,
		conn:   conn,
		send:   make(chan models.Message, 256),

		logger:   logger,
		auth:     auth,
		groupIDs: groupsIDs,
	}, nil
}

func (c *Client) readPump(h *Hub) {
	defer func() {
		h.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.logger.Error(fmt.Sprintf("error read message: %s", err))
			}
			break
		}

		log.Printf("Received message: %s \nfrom user with id: %d", message, c.userID)
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

		hasGroup := false
		for _, nGroupID := range message.GroupIDs {
			for _, uGroupID := range c.groupIDs {
				if uGroupID == nGroupID {
					hasGroup = true
				}
			}
		}

		if !hasGroup {
			continue
		}

		jmessage, err := json.Marshal(message)
		if err != nil {
			c.logger.Error(fmt.Sprintf("json marshal error message: %s", err))
			return
		}
		fmt.Println(c.conn.LocalAddr(), "send message")
		err = c.conn.WriteMessage(websocket.TextMessage, jmessage)
		if err != nil {
			c.logger.Error(fmt.Sprintf("write error:: %s", err))
			return
		}
	}
}

func (c *Client) testMessages() {
	count := 0
	for {
		count++
		time.Sleep(5 * time.Second)
		c.send <- models.Message{
			Header:   "Test",
			Text:     fmt.Sprintf("Test Notification: %d", count),
			GroupIDs: []uint{1},
		}
	}
}
