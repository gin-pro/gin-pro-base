package logx

import (
	"errors"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

func DefaultLog() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	return logger
}

func InitLogWithFile(paths PathMap) (*logrus.Logger, error) {
	if paths == nil || len(paths) == 0 {
		return nil, errors.New("LogLevel is empty, It has at least one")
	}
	logger := logrus.New()
	for level, p := range paths {
		logger.SetLevel(level)
		paths[level] = filepath.Join(p)
	}
	logger.AddHook(NewLfsHook(paths, &logrus.TextFormatter{}))
	return logger, nil
}
