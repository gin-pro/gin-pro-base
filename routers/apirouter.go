package routers

import "github.com/gin-gonic/gin"

type ApiRouter struct {
	engine *gin.Engine
}

func NewRouter() *ApiRouter {
	return &ApiRouter{
		engine: gin.Default(),
	}
}

func (c *ApiRouter) NewGroup(s string) *ApiGroup {
	return &ApiGroup{
		group: c.engine.Group(s),
	}
}
