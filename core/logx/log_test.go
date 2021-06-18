package logx

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLog(t *testing.T) {
	log := DefaultLog()
	log.Infof("i")
	log.Errorf("e")
}

func TestLogPath(t *testing.T) {
	l := map[logrus.Level]string{}
	l[logrus.ErrorLevel] = "./error.log"
	l[logrus.DebugLevel] = "./debug.log"
	l[logrus.InfoLevel] = "./info.log"
	log, _ := InitLogWithFile(l)
	log.Infof("i")
	log.Debug("e")
	log.Error("e")
}
