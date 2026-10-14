package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data"`
}

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Successfully",
		Data:    data,
	})
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "Successfully",
		Data:    data,
	})
}

func BadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, Response{
		Success: false,
		Code:    "bad_request",
		Message: err.Error(),
		Data:    nil,
	})
}

func NotFound(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusNotFound, Response{
		Success: false,
		Code:    "not_found",
		Message: err.Error(),
		Data:    nil,
	})
}

func UnprocessableEntity(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, Response{
		Success: false,
		Code:    "unprocessable_entity",
		Message: err.Error(),
		Data:    nil,
	})
}

func InternalServerError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
		Success: false,
		Code:    "internal_server_error",
		Message: err.Error(),
		Data:    nil,
	})
}
