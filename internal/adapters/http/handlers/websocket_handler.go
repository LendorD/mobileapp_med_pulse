package handlers

import (
	"net/http"
	"strconv"

	"github.com/AlexanderMorozov1919/mobileapp/internal/middleware/logging"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services/websocket"
	"github.com/gin-gonic/gin"
)

type WebsocketHandler struct {
	Handler *Handler
	Hub     *websocket.Hub
	logger  *logging.Logger
}

func NewWebsocketHandler(
	baseHandler *Handler,
	hub *websocket.Hub,
	logger *logging.Logger,
) *WebsocketHandler {
	return &WebsocketHandler{
		Handler: baseHandler,
		Hub:     hub,
		logger:  logger,
	}
}

// Register godoc
// @Summary Подписаться на уведомления
// @Tags Notification
// @Accept json
// @Produce json
// @Param user_id path int true "User id"
// @Success 200
// @Router /ws/notification/register/{user_id} [get]
func (ws *WebsocketHandler) Register(c *gin.Context) {
	strUserId := c.Param("user_id")
	if strUserId == "" {
		ws.Handler.ErrorResponse(c, nil, http.StatusBadRequest, "parameter 'user_id' must be exist", false)
		return
	}

	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		ws.Handler.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'user_id' must be an integer", false)
		return
	}

	err = ws.Hub.ServeRegister(c.Writer, c.Request, uint(userId))
	if err != nil {
		ws.Handler.ErrorResponse(c, err, http.StatusInternalServerError, "failed to register websocket", false)
		return
	}

	// WebSocket не возвращает JSON — соединение установлено
	// Ничего не отправляем — управление передано WebSocket
}

// Unregister godoc
// @Summary Отписаться от уведомлений
// @Tags Notification
// @Accept json
// @Produce json
// @Param user_id path int true "User id"
// @Success 200 {object} map[string]string
// @Router /ws/notification/unregister/{user_id} [get]
func (ws *WebsocketHandler) Unregister(c *gin.Context) {
	strUserId := c.Param("user_id")
	if strUserId == "" {
		ws.Handler.ErrorResponse(c, nil, http.StatusBadRequest, "parameter 'user_id' must be exist", false)
		return
	}

	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		ws.Handler.ErrorResponse(c, err, http.StatusBadRequest, "parameter 'user_id' must be an integer", false)
		return
	}

	ws.Hub.ServeUnregister(c.Writer, c.Request, uint(userId))
	ws.Handler.ResultResponse(c, "Success unregister notification subscriber", Empty, nil)
}
