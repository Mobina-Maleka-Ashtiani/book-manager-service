package Presentation

import (
	"book-manager-service/DataAccess"
	"github.com/sirupsen/logrus"
)

type BookManagerServer struct {
	Db     *DataAccess.GormDB
	Logger *logrus.Logger
}
