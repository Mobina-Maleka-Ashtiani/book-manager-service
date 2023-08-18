package Presentation

import (
	"book-manager-service/BusinessLogic"
	"book-manager-service/DataAccess"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (bms *BookManagerServer) HandleSignUp(context *gin.Context) {
	var user DataAccess.User
	if err := context.ShouldBindJSON(&user); err != nil {
		bms.Logger.WithError(err).Warn("con not read the request data and convert it to json")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "con not read the request data and convert it to json"})
		return
	}
	if err := SignUpInputValidation(user); err != nil {
		bms.Logger.WithError(err).Warn("invalid information")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := BusinessLogic.AddUser(bms.Db, user); err != nil {
		bms.Logger.WithError(err).Warn("can not create a new user")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "sign up successfully"})

}
