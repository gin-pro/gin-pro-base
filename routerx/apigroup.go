package routerx

import (
	"github.com/gin-gonic/gin"
)

type ApiGroup struct {
	group *gin.RouterGroup
}

func (r *ApiGroup) NewGroup(path string) *ApiGroup {
	return &ApiGroup{
		group: r.group.Group(path),
	}
}

func DefaultGroup(path string) *ApiGroup {
	return &ApiGroup{
		group: gin.Default().Group(path),
	}
}

func (r *ApiGroup) Use(middleware ...gin.HandlerFunc) *ApiGroup {
	r.group.Use(middleware...)
	return r
}

func (r *ApiGroup) Handle(httpMethod, relativePath string, fn interface{}) *ApiGroup {
	r.group.Handle(httpMethod, relativePath, RequestBind(fn))
	return r
}

func (r *ApiGroup) Any(relativePath string, fn interface{}) *ApiGroup {
	r.group.Any(relativePath, RequestBind(fn))
	return r
}
func (r *ApiGroup) POST(relativePath string, fn interface{}) *ApiGroup {
	r.group.POST(relativePath, RequestBind(fn))
	return r
}
func (r *ApiGroup) GET(relativePath string, fn interface{}) *ApiGroup {
	r.group.GET(relativePath, RequestBind(fn))
	return r
}

func (r *ApiGroup) DELETE(relativePath string, fn interface{}) *ApiGroup {
	r.group.DELETE(relativePath, RequestBind(fn))
	return r
}

func (r *ApiGroup) PATCH(relativePath string, fn interface{}) *ApiGroup {
	r.group.PATCH(relativePath, RequestBind(fn))
	return r
}

func (r *ApiGroup) PUT(relativePath string, fn interface{}) *ApiGroup {
	r.group.PUT(relativePath, RequestBind(fn))
	return r
}

func (r *ApiGroup) OPTIONS(relativePath string, fn interface{}) *ApiGroup {
	r.group.OPTIONS(relativePath, RequestBind(fn))
	return r
}

func (r *ApiGroup) HEAD(relativePath string, fn interface{}) *ApiGroup {
	r.group.HEAD(relativePath, RequestBind(fn))
	return r
}
