package logger

import log "github.com/sirupsen/logrus"

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdo)
}
