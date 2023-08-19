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
func (bms *BookManagerServer) HandleLogin(context *gin.Context) {
	var user DataAccess.User
	if err := context.ShouldBindJSON(&user); err != nil {
		bms.Logger.WithError(err).Warn("con not read the request data and convert it to json")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "con not read the request data and convert it to json"})
		return
	}

	if err := LoginInputValidation(user); err != nil {
		bms.Logger.WithError(err).Warn("invalid username or password")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := BusinessLogic.CheckUserCredential(bms.Db, user); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, err := BusinessLogic.GenerateJWTToken(user.Username)
	if err != nil {
		bms.Logger.WithError(err).Warn("token generation failed")
		context.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "token generation failed"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"access_token": token})
}

func (bms *BookManagerServer) HandleCreateBook(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")

	username, err := BusinessLogic.DecodeJWTToken(accessToken)
	if err != nil {
		bms.Logger.WithError(err).Warn("failed to decode access token")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to decode access token"})
		return
	}
	
	user, err := BusinessLogic.FindUserByUsername(bms.Db, username)
	if err != nil {
		bms.Logger.WithError(err).Warn("user not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "failed to decode access token"})
		return
	}

	var book DataAccess.Book
	if err := context.ShouldBindJSON(&book); err != nil {
		bms.Logger.WithError(err).Warn("con not read the request data and convert it to json")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "con not read the request data and convert it to json"})
		return
	}
}
