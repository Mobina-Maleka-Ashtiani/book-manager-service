package Presentation

import (
	"book-manager-service/DataAccess"
	"regexp"
	"strings"
)

func firstNameValidation(firstName string) bool {
	str := "abcdefghijklmnopqrtsuvwxyzABCDEFGEHIJKLMNOPQRSTUVWXYZ"
	if len(firstName) < 1 {
		return false
	}
	for i := 0; i < len(firstName); i++ {
		if !strings.Contains(str, string(firstName[i])) {
			return false
		}
	}
	return true
}

func lastNameValidation(lastName string) bool {
	str := "abcdefghijklmnopqrtsuvwxyzABCDEFGEHIJKLMNOPQRSTUVWXYZ"
	if len(lastName) < 1 {
		return false
	}
	for i := 0; i < len(lastName); i++ {
		if !strings.Contains(str, string(lastName[i])) {
			return false
		}
	}
	return true
}

func genderValidation(gender DataAccess.Gender) bool {
	if gender != DataAccess.Female && gender != DataAccess.Male && gender != DataAccess.NonBinary && gender != DataAccess.Transgender && gender != DataAccess.Intersex && gender != DataAccess.Other {
		return false
	}
	return true
}

func phoneNumberValidation(phoneNumber string) bool {
	if phoneNumber == "" {
		return false
	}
	pattern := `((\+|\(|0)?\d{1,3})?((\s|\)|\-))?(\d{10})$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(phoneNumber)

}

func usernameValidation(username string) bool {
	str := "abcdefghijklmnopqrtsuvwxyzABCDEFGEHIJKLMNOPQRSTUVWXYZ0123456789-_"
	if len(username) < 1 {
		return false
	}
	for i := 0; i < len(username); i++ {
		if !strings.Contains(str, string(username[i])) {
			return false
		}
	}
	return true
}

func emailValidation(email string) bool {
	pattern := `^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}
