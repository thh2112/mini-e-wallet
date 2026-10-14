package middleware

import (
	"mini-e-wallet/pkg/consts"
	"mini-e-wallet/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader(consts.HeaderXUserID)
		if userID != "" {
			c.Set(consts.CtxKeyUserID, userID)
			c.Next()
			return
		}

		authHeader := c.GetHeader(consts.HeaderAuth)
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenStr != "" {
				c.Set(consts.CtxKeyUserID, tokenStr)
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{
			Success: false,
			Code:    "unauthorized",
			Message: "unauthorized: missing or invalid user credentials",
			Data:    nil,
		})
	}
}
