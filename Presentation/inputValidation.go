package Presentation

import (
	"book-manager-service/DataAccess"
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
