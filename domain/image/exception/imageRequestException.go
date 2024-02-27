package exception

import (
	"clove-api/global/enum"
	"clove-api/global/exception"
	"clove-api/global/static"
	"fmt"
	"net/http"
)

type ImageRequestException struct {
	Code        int
	Message     string
	Description string
	ErrorData   string
}

func NewImageRequestException(err string) *ImageRequestException {
	return &ImageRequestException{
		Code:        http.StatusBadRequest,
		Message:     http.StatusText(http.StatusBadRequest),
		Description: "The image could not be parse.",
		ErrorData:   err,
	}
}

func (imagerequestException *ImageRequestException) Error() string {
	return fmt.Sprintf("[CUSTOM],%d,%s?%s", imagerequestException.Code, static.MessageFormatWithDescription(enum.ResponseType("FAILURE").Value(), imagerequestException.Message, imagerequestException.Description), imagerequestException.ErrorData)
}

func (imagerequestException *ImageRequestException) Throw() {
	exception.ThrowError(imagerequestException)
}
