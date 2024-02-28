package exception

import (
	"clove-api/global/enum"
	"clove-api/global/exception"
	"clove-api/global/static"
	"fmt"
	"net/http"
)

type PostRequestException struct {
	Code        int
	Message     string
	Description string
	ErrorData   string
}

func NewPostRequestException(err error) *PostRequestException {
	data := ""
	if err != nil {
		data = err.Error()
	}
	return &PostRequestException{
		Code:        http.StatusBadRequest,
		Message:     http.StatusText(http.StatusBadRequest),
		Description: "The post could not be parse.",
		ErrorData:   data,
	}
}

func (postRequestException *PostRequestException) Error() string {
	return fmt.Sprintf("[CUSTOM],%d,%s?%s", postRequestException.Code, static.MessageFormatWithDescription(enum.ResponseType("FAILURE").Value(), postRequestException.Message, postRequestException.Description), postRequestException.ErrorData)
}

func (postRequestException *PostRequestException) Throw() {
	exception.ThrowError(postRequestException)
}
