package exceptions

import "fmt"

var ModelName = "user"

type UserException struct {
	Message string
	Model   string
}

func (e UserException) Error() string {
	errMsg := fmt.Sprintf("[%v] Failure. Message : {%v}.", e.Model, e.Message)
	return errMsg
}

func NewUserException(msg string) *UserException {
	return &UserException{Message: msg, Model: ModelName}
}
