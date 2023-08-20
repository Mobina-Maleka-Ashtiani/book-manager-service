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
		bms.Logger.WithError(err).Warn("con not unmarshal the request data")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "con not unmarshal the request data"})
		return
	}

	if err := SignUpInputValidation(user); err != nil {
		bms.Logger.WithError(err).Warn("invalid information")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := BusinessLogic.AddUser(bms.Db, user); err != nil {
		bms.Logger.WithError(err).Warn("can not create a new user")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "sign up successfully"})

}

func (bms *BookManagerServer) HandleLogin(context *gin.Context) {
	var user DataAccess.User
	if err := context.ShouldBindJSON(&user); err != nil {
		bms.Logger.WithError(err).Warn("con not unmarshal the request data")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "con not unmarshal the request data"})
		return
	}

	if err := LoginInputValidation(user); err != nil {
		bms.Logger.WithError(err).Warn("invalid username or password")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := BusinessLogic.CheckUserCredential(bms.Db, user); err != nil {
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := BusinessLogic.GenerateJWTToken(user.Username)
	if err != nil {
		bms.Logger.WithError(err).Warn("failed to token generation")
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "failed to token generation"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"access_token": token})
}

func (bms *BookManagerServer) HandleCreateBook(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")

	username, err := BusinessLogic.DecodeJWTToken(accessToken)
	if err != nil {
		bms.Logger.WithError(err).Warn("failed to decode access token")
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "failed to decode access token"})
		return
	}

	user, err := BusinessLogic.FindUserByUsername(bms.Db, username)
	if err != nil {
		bms.Logger.WithError(err).Warn("user not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	var bookRequest BusinessLogic.BookRequestAndResponse
	if err := context.ShouldBindJSON(&bookRequest); err != nil {
		bms.Logger.WithError(err).Warn("con not unmarshal the request data")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "con not unmarshal the request data"})
		return
	}

	if err := BusinessLogic.AddBookToUser(bms.Db, *user, bookRequest); err != nil {
		bms.Logger.WithError(err).Warn("failed to add book")
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "successfully added this book to your books"})
}

func (bms *BookManagerServer) HandleGetAllBooks(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")

	username, err := BusinessLogic.DecodeJWTToken(accessToken)
	if err != nil {
		bms.Logger.WithError(err).Warn("failed to decode access token")
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "failed to decode access token"})
		return
	}

	_, err = BusinessLogic.FindUserByUsername(bms.Db, username)
	if err != nil {
		bms.Logger.WithError(err).Warn("user not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
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
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "failed to decode access token"})
		return
	}

	_, err = BusinessLogic.FindUserByUsername(bms.Db, username)
	if err != nil {
		bms.Logger.WithError(err).Warn("user not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	bookResponse, err := BusinessLogic.GetBookByID(bms.Db, id)
	if err != nil {
		bms.Logger.WithError(err).Warn("failed to get book")
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, bookResponse)
}

func (bms *BookManagerServer) HandleUpdateBook(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")

	username, err := BusinessLogic.DecodeJWTToken(accessToken)
	if err != nil {
		bms.Logger.WithError(err).Warn("failed to decode access token")
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "failed to decode access token"})
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
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	userBook, err := BusinessLogic.FindUserBook(bms.Db, *user, id)
	if err != nil {
		bms.Logger.WithError(err).Warn("book not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var bur BusinessLogic.BookUpdateRequest
	if err := context.ShouldBindJSON(&bur); err != nil {
		bms.Logger.WithError(err).Warn("con not unmarshal the request data")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "con not unmarshal the request data"})
		return
	}

	if err := BusinessLogic.UpdateBook(bms.Db, *userBook, bur); err != nil {
		bms.Logger.WithError(err).Warn("failed to update book")
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "book updated successfully"})
}

func (bms *BookManagerServer) HandleDeleteBook(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")

	username, err := BusinessLogic.DecodeJWTToken(accessToken)
	if err != nil {
		bms.Logger.WithError(err).Warn("failed to decode access token")
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "failed to decode access token"})
		return
	}

	user, err := BusinessLogic.FindUserByUsername(bms.Db, username)
	if err != nil {
		bms.Logger.WithError(err).Warn("user not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	userBook, err := BusinessLogic.FindUserBook(bms.Db, *user, id)
	if err != nil {
		bms.Logger.WithError(err).Warn("book not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := BusinessLogic.DeleteBook(bms.Db, *userBook); err != nil {
		bms.Logger.WithError(err).Warn("failed to delete book")
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}
