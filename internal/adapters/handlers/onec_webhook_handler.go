package handlers

import (
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/gin-gonic/gin"
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
