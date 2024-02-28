package exception

import (
	"clove-api/global/enum"
	"clove-api/global/exception"
	"clove-api/global/static"
	"fmt"
	"net/http"
)

type ImageNotFoundException struct {
	Code        int
	Message     string
	Description string
	ErrorData   string
}

func NewImageNotFoundException(err error) *ImageNotFoundException {
	data := ""
	if err != nil {
		data = err.Error()
	}
	return &ImageNotFoundException{
		Code:        http.StatusNotFound,
		Message:     http.StatusText(http.StatusNotFound),
		Description: "The image could not be found.",
		ErrorData:   data,
	}
}

func (imageNotFoundException *ImageNotFoundException) Error() string {
	return fmt.Sprintf("[CUSTOM],%d,%s?%s", imageNotFoundException.Code, static.MessageFormatWithDescription(enum.ResponseType("FAILURE").Value(), imageNotFoundException.Message, imageNotFoundException.Description), imageNotFoundException.ErrorData)
}

func (imageNotFoundException *ImageNotFoundException) Throw() {
	exception.ThrowError(imageNotFoundException)
}
