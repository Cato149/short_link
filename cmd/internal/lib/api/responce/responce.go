package responce

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type Response struct {
	Status int    `json:"status"`
	Error  string `json:"error,omitempty"` //omitempty - параметер structTag JSON. IF Param empty -> Result JSON exclude Param
}

func OK() Response {
	return Response{
		Status: http.StatusOK,
	}
}

func ServerError(msg string) Response {
	return Response{
		Status: http.StatusInternalServerError,
		Error:  msg,
	}
}

func BadRequestError(msg string) Response {
	return Response{
		Status: http.StatusBadRequest,
		Error:  msg,
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("Field required: %s", err.Field()))
		case "url":
			errMsgs = append(errMsgs, fmt.Sprintf("Invalid URL: %s", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("Invalid field: %s", err.Field()))
		}
	}

	return Response{
		Status: http.StatusBadRequest,
		Error:  strings.Join(errMsgs, ". \n"),
	}
}
