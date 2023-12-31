package main

import (
	"book-manager-service/DataAccess"
	"book-manager-service/Presentation"
	"github.com/gin-gonic/gin"
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

	bookManagerServer := Presentation.BookManagerServer{
		Db:     gormDB,
		Logger: logger,
	}

	router := gin.Default()
	router.Use(Presentation.FirstGinMiddleware())
	router.Use(Presentation.SecondGinMiddleware())
	router.POST("/api/v1/auth/signup", bookManagerServer.HandleSignUp)
	router.POST("/api/v1/auth/login", bookManagerServer.HandleLogin)
	router.POST("/api/v1/books", bookManagerServer.HandleCreateBook)
	router.GET("/api/v1/books", bookManagerServer.HandleGetAllBooks)
	router.GET("/api/v1/books/:id", bookManagerServer.HandleGetBook)
	router.PUT("/api/v1/books/:id", bookManagerServer.HandleUpdateBook)
	router.DELETE("/api/v1/books/:id", bookManagerServer.HandleDeleteBook)

	router.Run("0.0.0.0:8080")
}
