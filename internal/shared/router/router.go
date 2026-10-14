package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RouterSet = wire.NewSet(
	NewRouterGroup,
)

type RouterGroup struct {
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{}
}

func (r *RouterGroup) Register(public, protected *gin.RouterGroup) {

}
