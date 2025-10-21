package handlers

import (
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// OneCWebhook godoc
// @Summary Webhook from 1C: new receptions
// @Tags 1C
// @Accept json
// @Produce json
// @Param update body models.Call true "Receptions update"
// @Success 200
// @Router /webhook/onec/receptions [post]
func (h *Handler) OneCWebhook(c *gin.Context) {
	var update models.Call
	if err := c.ShouldBindJSON(&update); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.usecase.HandleReceptionsUpdate(c.Request.Context(), update)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type AuthCreditionals struct {
	Users []OneCUser
}
type OneCUser struct {
	Login    string `json:"login" example:"doc1"`
	Password string `json:"password" example:"secret123"`
}

// OneCAuthWebhook receives a list of users from 1C and syncs them into the system.
// @Summary Sync users from 1C
// @Description Syncs a batch of users (login + password) received from 1C into the internal auth system.
// @Tags Webhooks
// @Accept json
// @Produce json
// @Param request body AuthCreditionals true "List of users to sync"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /onec/auth [post]
func (h *Handler) OneCAuthWebhook(c *gin.Context) {
	var update AuthCreditionals

	if err := c.ShouldBindJSON(&update); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	var users []entities.AuthUser
	for _, u := range update.Users {
		hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		users = append(users, entities.AuthUser{
			Login:    u.Login,
			Password: string(hash),
		})
	}

	if err := h.usecase.SyncUsers(c.Request.Context(), users); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "sync failed"})
		return
	}

	c.Status(http.StatusOK)
}
