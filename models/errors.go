package models

import "strings"

const (
	ErrNotFound          modelError   = "models: resource not found"
	ErrTitleRequired     modelError   = "models: title is required"
	ErrPasswordIncorrect modelError   = "models: incorrect password provided"
	ErrEmailRequired     modelError   = "models: email address is required"
	ErrEmailInvalid      modelError   = "models: email address is not valid"
	ErrEmailTaken        modelError   = "models: email address is already taken"
	ErrPasswordToShort   modelError   = "models: password at leat must be 8 characters long"
	ErrPasswordRequired  modelError   = "models: password required"
	ErrRememberTooShort  privateError = "models: remember token must be at least 32 bytes"
	ErrRememberRequired  privateError = "models: remember is required"
	ErrIDInvalid         privateError = "models: ID provided was invalid"
	ErrUserIDRequired    privateError = "models: user ID is required"
)

type modelError string
type privateError string

func (e modelError) Error() string {
	return string(e)
}

func (e privateError) Error() string {
	return string(e)
}

func (e modelError) Public() string {
	s := strings.Replace(string(e), "models: ", "", 1)
	split := strings.Split(s, " ")
	split[0] = strings.Title(split[0])
	return strings.Join(split, " ")
}
