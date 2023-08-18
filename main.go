package main

import (
	"book-manager-service/DataAccess"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

func main() {
	var cfg DataAccess.Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err.Error())
	}

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	gormDB, err := DataAccess.NewGormDB(cfg)
	if err != nil {
		logger.WithError(err).Fatal("error in connecting to the postgres database")
	}
	logger.Info("connected to the database")

	err = gormDB.CreateSchemas()
	if err != nil {
		logger.WithError(err).Fatal("error in database migration")
	}
	logger.Infoln("migrate tables and models successfully")

}
