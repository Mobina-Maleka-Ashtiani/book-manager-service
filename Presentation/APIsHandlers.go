package Presentation

import (
	"book-manager-service/DataAccess"
	"book-manager-service/Logic"
	"github.com/gin-gonic/gin"
	"net/http"
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

	if err := Logic.AddUser(bms.Db, user); err != nil {
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

	if err := Logic.CheckUserCredential(bms.Db, user); err != nil {
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := Logic.GenerateJWTToken(user.Username)
	if err != nil {
		bms.Logger.WithError(err).Warn("failed to token generation")
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "failed to token generation"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"access_token": token})
}

func (bms *BookManagerServer) HandleCreateBook(context *gin.Context) {
	un, ok := context.Get("username")
	if !ok {
		bms.Logger.Warn("unauthorized")
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized token"})
		return
	}
	username := un.(string)

	user, err := Logic.FindUserByUsername(bms.Db, username)
	if err != nil {
		bms.Logger.WithError(err).Warn("user not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	var bookRequest Logic.BookRequestAndResponse
	if err := context.ShouldBindJSON(&bookRequest); err != nil {
		bms.Logger.WithError(err).Warn("con not unmarshal the request data")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "con not unmarshal the request data"})
		return
	}

	if err := Logic.AddBookToUser(bms.Db, *user, bookRequest); err != nil {
		bms.Logger.WithError(err).Warn("failed to add book")
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "successfully added this book to your books"})
}

func (bms *BookManagerServer) HandleGetAllBooks(context *gin.Context) {
	un, ok := context.Get("username")
	if !ok {
		bms.Logger.Warn("unauthorized")
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized token"})
		return
	}
	username := un.(string)

	_, err := Logic.FindUserByUsername(bms.Db, username)
	if err != nil {
		bms.Logger.WithError(err).Warn("user not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	booksResponse, err := Logic.GetAllBooks(bms.Db)
	if err != nil {
		bms.Logger.WithError(err).Warn("failed to get books")
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to get books"})
		return
	}

	context.IndentedJSON(http.StatusOK, map[string]interface{}{"books": booksResponse})
}

func (bms *BookManagerServer) HandleGetBook(context *gin.Context) {
	un, ok := context.Get("username")
	if !ok {
		bms.Logger.Warn("unauthorized")
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized token"})
		return
	}
	username := un.(string)

	_, err := Logic.FindUserByUsername(bms.Db, username)
	if err != nil {
		bms.Logger.WithError(err).Warn("user not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	bi, ok := context.Get("bookId")
	if !ok {
		bms.Logger.Warn("invalid id")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}
	id := bi.(int)

	bookResponse, err := Logic.GetBookByID(bms.Db, id)
	if err != nil {
		bms.Logger.WithError(err).Warn("failed to get book")
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, bookResponse)
}

func (bms *BookManagerServer) HandleUpdateBook(context *gin.Context) {
	un, ok := context.Get("username")
	if !ok {
		bms.Logger.Warn("unauthorized")
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized token"})
		return
	}
	username := un.(string)

	user, err := Logic.FindUserByUsername(bms.Db, username)
	if err != nil {
		bms.Logger.WithError(err).Warn("user not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	bi, ok := context.Get("bookId")
	if !ok {
		bms.Logger.Warn("invalid id")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}
	id := bi.(int)

	userBook, err := Logic.FindUserBook(bms.Db, *user, id)
	if err != nil {
		bms.Logger.WithError(err).Warn("book not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var bur Logic.BookUpdateRequest
	if err := context.ShouldBindJSON(&bur); err != nil {
		bms.Logger.WithError(err).Warn("con not unmarshal the request data")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "con not unmarshal the request data"})
		return
	}

	if err := Logic.UpdateBook(bms.Db, *userBook, bur); err != nil {
		bms.Logger.WithError(err).Warn("failed to update book")
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "book updated successfully"})
}

func (bms *BookManagerServer) HandleDeleteBook(context *gin.Context) {
	un, ok := context.Get("username")
	if !ok {
		bms.Logger.Warn("unauthorized")
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized token"})
		return
	}
	username := un.(string)

	user, err := Logic.FindUserByUsername(bms.Db, username)
	if err != nil {
		bms.Logger.WithError(err).Warn("user not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	bi, ok := context.Get("bookId")
	if !ok {
		bms.Logger.Warn("invalid id")
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	}
	id := bi.(int)

	userBook, err := Logic.FindUserBook(bms.Db, *user, id)
	if err != nil {
		bms.Logger.WithError(err).Warn("book not found")
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := Logic.DeleteBook(bms.Db, *userBook); err != nil {
		bms.Logger.WithError(err).Warn("failed to delete book")
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted successfully"})
}
