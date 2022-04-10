package validator

import (
	"regexp"
	"unicode"
)

func IsEmail(email string) bool {
	r := regexp.MustCompile(`^[a-zA-Z0-9+-\_.]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)
	return r.MatchString(email)
}

/*
*	Dev By Kyle
*	대소문자, 특수문자 각각 1개 이상, n자 이상
 */
func IsRightPassword(password string, n int) bool {
	var (
		hasMinLen  = false //최소 길이
		hasUpper   = false //최소 대문자 1개
		hasLower   = false //최소 소문자 1개
		hasNumber  = false //최소 숫자 1개
		hasSpecial = false //최소 특수문자 1개
	)
	if len(password) >= n {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
