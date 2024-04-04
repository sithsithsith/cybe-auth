package exceptions

import "fmt"

type ServiceError struct {
	Reason string
}

func (s ServiceError) Error() string {
	errMsg := fmt.Sprintf("Service Error Ocurred. \n Reason : {%v}.", s.Reason)
	return errMsg
}

func NewServiceError(r string) *ServiceError {
	return &ServiceError{Reason: r}
}
