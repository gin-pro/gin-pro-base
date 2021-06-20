package routerx

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestNew(t *testing.T) {
	router := NewRouter(nil)

	group := router.NewGroup("/hello")
	group.POST("test", Hello)

}

func Hello(c *gin.Context, i interface{}) {

}
