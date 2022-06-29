package commonx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	msgs   []interface{}
	lc     string
	status int
}

func NewResponse(code int) *Response {
	return &Response{Code: code}
}

func NewResOk(data ...interface{}) *Response {
	c := &Response{Code: ResponseCodeOk}
	if len(data) > 0 {
		c.Data = data[0]
	}
	return c
}

func (c *Response) SetCode(code int) *Response {
	c.Code = code
	return c
}
func (c *Response) SetStatus(a int) *Response {
	c.status = a
	return c
}
func (c *Response) SetMsgf(format string, args ...interface{}) *Response {
	c.Msg = fmt.Sprintf(format, args...)
	return c
}
func (c *Response) SetLc(lc string) *Response {
	c.lc = lc
	return c
}
func (c *Response) SetMsgObjs(args ...interface{}) *Response {
	c.msgs = append(c.msgs, args...)
	return c
}
func (c *Response) SetData(data interface{}) *Response {
	c.Data = data
	return c
}
func (c *Response) ResJson(g *gin.Context, format ...bool) {
	if c.Msg == "" {
		if c.lc == "" {
			c.lc = "zh" //默认
		}
		c.Msg = GetResponseMsg(c.lc, c.Code, c.msgs...)
	}
	stat := http.StatusOK
	if c.status != 0 {
		stat = c.status
	}
	if len(format) > 0 && format[0] {
		g.IndentedJSON(stat, c)
	} else {
		g.JSON(stat, c)
	}
}
