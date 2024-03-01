package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const apiKeyAuth = "7h1sI$ApiK3y"

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h *Handlers) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "empty auth header",
		})
		c.Abort()
		return
	}

	if header != apiKeyAuth {
		c.JSON(http.StatusUnauthorized, HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid auth header",
		})
		c.Abort()
		return
	}
	c.Next()
}
