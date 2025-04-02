package logger

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
}

func Info(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Info(msg)
}

func Error(msg string, fields map[string]interface{}) {
	log.WithFields(fields).Error(msg)
}