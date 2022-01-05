package ginpro

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"runtime/debug"
)

func Recovers(pc uintptr, file string, line int, ok bool, fn ...func(interface{})) {
	if err := recover(); err != nil {
		if len(fn) > 0 {
			fn[0](err)
			return
		}
		if ok {
			logrus.Warnf("recover %s [%d]", file, line)
			logrus.Warnf("err:%v", err)
			logrus.Warnf("Stack:%s", string(debug.Stack()))
		} else {
			logrus.Warnf("recover:%v", err)
			logrus.Warnf("Stack:%s", string(debug.Stack()))
		}
	}
}

func Go(fn func()) {
	if fn == nil {
		return
	}
	go func() {
		defer Recovers(runtime.Caller(1))
		fn()
	}()
}
