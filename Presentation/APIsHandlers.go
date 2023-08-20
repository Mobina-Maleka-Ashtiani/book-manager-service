package Presentation

import (
	"book-manager-service/BusinessLogic"
	"book-manager-service/DataAccess"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	var bookRequest BusinessLogic.BookRequestAndResponse
	if err := context.ShouldBindJSON(&bookRequest); err != nil {
		bms.Logger.WithError(err).Warn("con not read the request data and convert it to json")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "con not read the request data and convert it to json"})
		return
	}

	if err := BusinessLogic.AddBookToUser(bms.Db, *user, bookRequest); err != nil {
		bms.Logger.WithError(err).Warn("failed to add book")
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "successfully added this book to your books"})
}

func (bms *BookManagerServer) HandleGetAllBooks(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")

	username, err := BusinessLogic.DecodeJWTToken(accessToken)
	if err != nil {
		bms.Logger.WithError(err).Warn("failed to decode access token")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to decode access token"})
		return
	}

	_, err = BusinessLogic.FindUserByUsername(bms.Db, username)
	if err != nil {
		bms.Logger.WithError(err).Warn("user not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	booksResponse, err := BusinessLogic.GetAllBooks(bms.Db)
	if err != nil {
		bms.Logger.WithError(err).Warn("failed to get books")
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to get books"})
		return
	}

	context.IndentedJSON(http.StatusOK, map[string]interface{}{"books": booksResponse})
}

func (bms *BookManagerServer) HandleGetBook(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")

	username, err := BusinessLogic.DecodeJWTToken(accessToken)
	if err != nil {
		bms.Logger.WithError(err).Warn("failed to decode access token")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to decode access token"})
		return
	}

	_, err = BusinessLogic.FindUserByUsername(bms.Db, username)
	if err != nil {
		bms.Logger.WithError(err).Warn("user not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	bookResponse, err := BusinessLogic.GetBookByID(bms.Db, id)
	if err != nil {
		bms.Logger.WithError(err).Warn("failed to get book")
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, bookResponse)
}

func (bms *BookManagerServer) HandleUpdateBook(context *gin.Context) {
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
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	userBook, err := BusinessLogic.FindUserBook(bms.Db, *user, id)
	if err != nil {
		bms.Logger.WithError(err).Warn("book not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

}
