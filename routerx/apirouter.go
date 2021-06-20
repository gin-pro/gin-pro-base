package routerx

import "github.com/gin-gonic/gin"

type ApiEngine struct {
	engine *gin.Engine
}

func NewRouter(engine *gin.Engine) *ApiEngine {
	return &ApiEngine{
		engine: engine,
	}
}

func (c *ApiEngine) NewGroup(s string) *ApiGroup {
	return &ApiGroup{
		group: c.engine.Group(s),
	}
}
func (c *ApiEngine) Run(addr ...string) error {
	return c.engine.Run(addr...)
}

func (r *ApiEngine) Use(middleware ...gin.HandlerFunc) *ApiEngine {
	r.engine.Use(middleware...)
	return r
}
