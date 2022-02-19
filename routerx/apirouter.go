package routerx

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

type ApiEngine struct {
	engine *gin.Engine
}

func Default() *ApiEngine {
	return &ApiEngine{
		engine: gin.Default(),
	}
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

func (r *ApiEngine) Engine() *gin.Engine {
	r.engine.LoadHTMLFiles()
	return r.engine
}

func (r *ApiEngine) LoadHTMLFiles(files ...string) {
	r.engine.LoadHTMLFiles(files...)
}

func (r *ApiEngine) SetHTMLTemplate(templ *template.Template) {
	r.engine.SetHTMLTemplate(templ)
}

func (r *ApiEngine) LoadHTMLGlob(pattern string) {
	r.engine.LoadHTMLGlob(pattern)
}
