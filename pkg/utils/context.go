package utils

import (
	"mini-e-wallet/pkg/consts"

	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) string {
	if userID := c.GetString(consts.CtxKeyUserID); userID != "" {
		return userID
	}
	return c.GetHeader(consts.HeaderXUserID)
}

func GetRequestID(c *gin.Context) string {
	if reqID := c.GetString(consts.CtxKeyRequestID); reqID != "" {
		return reqID
	}
	return c.GetHeader(consts.HeaderXRequestID)
}
