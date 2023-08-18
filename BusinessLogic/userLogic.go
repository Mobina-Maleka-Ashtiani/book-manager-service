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
		//log.Println(err, " password not hash")
		return err
	} else {
		user.Password = string(pw)
	}
}
