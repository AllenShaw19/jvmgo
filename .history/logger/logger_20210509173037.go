package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// 从配置文件中获取
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
