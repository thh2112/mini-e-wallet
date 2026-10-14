package middleware

import (
	"mini-e-wallet/pkg/consts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader(consts.HeaderXRequestID)
		if reqID == "" {
			reqID = uuid.New().String()
		}

		c.Set(consts.CtxKeyRequestID, reqID)
		c.Writer.Header().Set(consts.HeaderXRequestID, reqID)
		c.Next()
	}
}
