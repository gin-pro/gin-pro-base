package runtimex

import (
	"github.com/sirupsen/logrus"
)

func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}
	if p := recover(); p != nil {
		logrus.Warning(p)
	}
}

func Go(fn func()) {
	if fn == nil {
		return
	}
	go func() {
		defer Recover()
		fn()
	}()
}
