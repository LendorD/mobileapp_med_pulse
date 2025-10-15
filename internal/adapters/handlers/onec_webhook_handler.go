package handlers

import (
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// OneCWebhook godoc
// @Summary Webhook from 1C: new receptions
// @Tags 1C
// @Accept json
// @Produce json
// @Param update body models.OneCReceptionsUpdate true "Receptions update"
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

func (h *Handler) OneCAuthWebhook(c *gin.Context) {
	var update struct {
		Users []struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		} `json:"users"`
	}

	if err := c.ShouldBindJSON(&update); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	var users []sqlite.AuthUser
	for _, u := range update.Users {
		hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		users = append(users, sqlite.AuthUser{
			Login:    u.Login,
			Password: string(hash),
		})
	}

	if err := h.authUsecase.SyncUsers(c.Request.Context(), users); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "sync failed"})
		return
	}

	c.Status(http.StatusOK)
}
