package Presentation

import "strings"

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
