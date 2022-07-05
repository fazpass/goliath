package logger

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Debug    bool
	Channel  string
	Path     string
	Filename string
}

func InitLoger(config Config) *logrus.Logger {
	var logger = logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	if !config.Debug {
		logger.Out = ioutil.Discard
	}

	switch config.Channel {
	case "file":
		var path = config.Path
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.MkdirAll(path, 0777)
			if err != nil {
				logger.Fatal(err)
			}
		}

		var filePath = path + "/" + config.Filename
		var file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logger.Fatal(err)
		}
		logger.SetOutput(file)
	}

	return logger
}
