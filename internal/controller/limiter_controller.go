package controller

import (
	"net/http"

	"github.com/seyedmo30/http_request_limiter/internal/interfaces"

	"github.com/gin-gonic/gin"
)

type LimiterController struct {
	service interfaces.LimiterService
}

func NewLimiterController(service interfaces.LimiterService) *LimiterController {
	return &LimiterController{service: service}
}

func (c *LimiterController) HandleRequest(ctx *gin.Context) {
	clientID := ctx.Query("client_id")
	if clientID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "client_id is required"})
		return
	}

	userAllowed, globalAllowed := c.service.HandleRequest(clientID)
	if !userAllowed {
		ctx.JSON(http.StatusTooManyRequests, gin.H{"status": "user limit exceeded"})
		return
	}

	if !globalAllowed {
		ctx.JSON(http.StatusTooManyRequests, gin.H{"status": "global traffic limit exceeded"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "allowed"})
}
