package main

import (
	"lenta-parser/config"

	"github.com/sirupsen/logrus"
)

func main(){
	logger := logrus.New()
    logger.SetFormatter(&logrus.JSONFormatter{})

	cnf, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info(cnf)
}