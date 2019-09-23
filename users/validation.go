package users

import (
	"errors"
	"regexp"
)

//Password
const usersPasswordLength int = 16

//Email
const usersEmailMinLength int = 8
const usersEmailMaxLength int = 255

//Username
const usersUserameMinLength int = 3
const usersUsernameMaxLength int = 32

//IsValidUsername username validate
func IsValidUsername(username string) (result bool, err error) {
	if len(username) >= usersUsernameMaxLength || len(username) <= usersUserameMinLength {
		return
	}
	result, err = regexp.MatchString(`^[a-z]+[a-z0-9]+([a-z0-9](\_|\.)[a-z0-9])*[a-z0-9]+$`, username)
	return
}

//IsValidEmail email validate
func IsValidEmail(email string) (result bool, err error) {
	if len(email) <= usersEmailMinLength || len(email) >= usersEmailMaxLength {
		return
	}
	result, err = regexp.MatchString(`^[A-Za-z0-9]+(\.|\_|[A-Za-z0-9])[A-Za-z-0-9]+@[A-Za-z]+(\.[A-Za-z0-9]+){1,}`, email)
	return
}

//IsValidPassword password validate
func IsValidPassword(password string) (result bool, err error) {
	result, err = IsHexString(password)
	if result == false || err != nil {
		//Make sure it false
		result = false
		return
	}
	result = len(password) >= usersPasswordLength
	return
}

//IsHexString hex string validate
func IsHexString(hexString string) (result bool, err error) {
	if len(hexString)%2 == 1 {
		err = errors.New("Invalid hex string")
		return
	}
	result, err = regexp.MatchString(`^[A-Fa-f0-9]{1,}$`, hexString)
	return
}

//IsValidField check field name valid
func IsValidField(field string) (result bool, err error) {
	validMap := map[string]bool{
		"first-name":  true,
		"last-name":   true,
		"address":     true,
		"id-number":   true,
		"issued-date": true}
	result = validMap[field]
	return
}

//IsValidValue validate input string
func IsValidValue(value string) (result bool, err error) {
	result, err = regexp.MatchString(`^[^\x00^\x07^\x09^\x0a^\x0d^\x22^\x27^\x60^“^”^‘^’]+$`, value)
	return
}
