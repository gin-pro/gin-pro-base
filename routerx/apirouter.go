package routerx

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
func (c *ApiRouter) Run(addr ...string) error {
	return c.engine.Run(addr...)
}
