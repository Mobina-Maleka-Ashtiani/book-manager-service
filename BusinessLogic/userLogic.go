package BusinessLogic

import (
	"book-manager-service/DataAccess"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(gdb *DataAccess.GormDB, user DataAccess.User) error {
	if gdb.UsernameExistence(user.Username) {
		return errors.New("this username already taken")
	}

	if gdb.EmailExistence(user.Email) {
		return errors.New("this email already taken")
	}

	if gdb.PhoneNumberExistence(user.PhoneNumber) {
		return errors.New("this phone number already taken")
	}

	if pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 0); err != nil {
		return err
	} else {
		user.Password = string(pw)
	}
	return gdb.AddUserToDatabase(&user)
}

func CheckUserCredential(gdb *DataAccess.GormDB, user DataAccess.User) error {
	foundUser, err := gdb.FindUserByUsername(user.Username)
	if err != nil {
		return errors.New("there is no user with this username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		return errors.New("the password is incorrect")
	}

	return nil
}
