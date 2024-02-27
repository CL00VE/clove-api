package exception

import (
	"clove-api/global/enum"
	"clove-api/global/exception"
	"clove-api/global/static"
	"fmt"
	"net/http"
)

type ProfileNotFoundException struct {
	Code        int
	Message     string
	Description string
	ErrorData   string
}

func NewProfileNotFoundException(err string) *ProfileNotFoundException {
	return &ProfileNotFoundException{
		Code:        http.StatusNotFound,
		Message:     http.StatusText(http.StatusNotFound),
		Description: "The profile could not be found.",
		ErrorData:   err,
	}
}

func (profileNotFoundException *ProfileNotFoundException) Error() string {
	return fmt.Sprintf("[CUSTOM],%d,%s?%s", profileNotFoundException.Code, static.MessageFormatWithDescription(enum.ResponseType("FAILURE").Value(), profileNotFoundException.Message, profileNotFoundException.Description), profileNotFoundException.ErrorData)
}

func (profileNotFoundException *ProfileNotFoundException) Throw() {
	exception.ThrowError(profileNotFoundException)
}
