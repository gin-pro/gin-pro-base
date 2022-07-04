package routerx

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

func RequestBind(fn interface{}) gin.HandlerFunc {
	return requestBind(fn)
}

func requestBind(fn interface{}) gin.HandlerFunc {
	fnv := reflect.ValueOf(fn)
	if fnv.Kind() != reflect.Func {
		return nil
	}
	fnt := fnv.Type()
	return func(c *gin.Context) {
		inls := bind(c, fnt)
		defer func() {
			if err := recover(); err != nil {
				c.String(500, fmt.Sprintf("router errs:%+v", err))
			}
		}()
		fnv.Call(inls)
	}
}

func json(c *gin.Context, argtr reflect.Type) (reflect.Value, error) {
	if argtr.Kind() == reflect.Struct || argtr.Kind() == reflect.Map {
		argv := reflect.New(argtr)
		if err := c.BindJSON(argv.Interface()); err != nil {
			return argv, err
		}
		if argv.Kind() != reflect.Ptr {
			return argv.Elem(), nil
		}
		return argv, nil
	}
	return reflect.Value{}, errors.New("argtr not Struct or Map")
}

func from(c *gin.Context, argtr reflect.Type) (reflect.Value, error) {
	if argtr.Kind() == reflect.Struct || argtr.Kind() == reflect.Map {
		argv := reflect.New(argtr)
		if err := c.Bind(argv.Interface()); err != nil {
			return argv, err
		}
		if argv.Kind() != reflect.Ptr {
			return argv.Elem(), nil
		}
		return argv, nil
	}
	return reflect.Value{}, errors.New("argtr not Struct or Map")
}

func bind(c *gin.Context, fnt reflect.Type) []reflect.Value {
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
			argv, err := json(c, argtr)
			if err != nil {
				c.String(500, fmt.Sprintf("params err[%d]:%+v", i, err))
				return nil
			}
			inls[i] = argv
		} else if strings.Contains(c.ContentType(), "multipart/form-data") {
			argv, err := from(c, argtr)
			if err != nil {
				c.String(500, fmt.Sprintf("params err[%d]:%+v", i, err))
				return nil
			}
			inls[i] = argv
		} else {
			argv, err := json(c, argtr)
			if err != nil {
				c.String(500, fmt.Sprintf("params err[%d]:%+v", i, err))
				return nil
			}
			inls[i] = argv
		}
	}
	return inls
}
