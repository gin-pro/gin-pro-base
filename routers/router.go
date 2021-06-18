package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

func RequestParseJson(fn interface{}) gin.HandlerFunc {
	return requestParseJson(fn)
}

func requestParseJson(fn interface{}) gin.HandlerFunc {
	fnv := reflect.ValueOf(fn)
	if fnv.Kind() != reflect.Func {
		return nil
	}
	fnt := fnv.Type()
	return func(c *gin.Context) {
		nmIn := fnt.NumIn()
		inls := make([]reflect.Value, nmIn)
		inls[0] = reflect.ValueOf(c)
		for i := 1; i < nmIn; i++ {
			argt := fnt.In(i)
			argtr := argt
			if argt.Kind() == reflect.Ptr {
				argtr = argt.Elem()
			}
			inls[i] = reflect.Zero(argt)
			if strings.Contains(c.ContentType(), "application/json") {
				if argtr.Kind() == reflect.Struct || argtr.Kind() == reflect.Map {
					argv := reflect.New(argtr)
					if err := c.BindJSON(argv.Interface()); err != nil {
						c.String(500, fmt.Sprintf("params err[%d]:%+v", i, err))
						return
					}
					if argt.Kind() == reflect.Ptr {
						inls[i] = argv
					} else {
						inls[i] = argv.Elem()
					}
				}
			}
		}
		defer func() {
			if err := recover(); err != nil {
				c.String(500, fmt.Sprintf("router errs:%+v", err))
			}
		}()
		fnv.Call(inls)
	}
}
