package exception

import (
	"clove-api/global/enum"
	"clove-api/global/exception"
	"clove-api/global/static"
	"fmt"
	"net/http"
)

type PostNotFoundException struct {
	Code        int
	Message     string
	Description string
	ErrorData   string
}

func NewPostNotFoundException(err string) *PostNotFoundException {
	return &PostNotFoundException{
		Code:        http.StatusNotFound,
		Message:     http.StatusText(http.StatusNotFound),
		Description: "The post could not be found.",
		ErrorData:   err,
	}
}

func (postNotFoundException *PostNotFoundException) Error() string {
	return fmt.Sprintf("[CUSTOM],%d,%s?%s", postNotFoundException.Code, static.MessageFormatWithDescription(enum.ResponseType("FAILURE").Value(), postNotFoundException.Message, postNotFoundException.Description), postNotFoundException.ErrorData)
}

func (postNotFoundException *PostNotFoundException) Throw() {
	exception.ThrowError(postNotFoundException)
}
