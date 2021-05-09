package logger

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// TODO: 从配置文件中获取
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func Debug(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Debug(msg)
}

func Warn(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Warn(msg)
}

func Info(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Info(msg)
}

func Error(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Error(msg)
}

func Fatal(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Debug(msg)
}

