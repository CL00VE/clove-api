package exception

import (
	"clove-api/global/enum"
	"clove-api/global/exception"
	"clove-api/global/static"
	"fmt"
	"net/http"
)

type ProfileRequestException struct {
	Code        int
	Message     string
	Description string
	ErrorData   string
}

func NewProfileRequestException(err string) *ProfileRequestException {
	return &ProfileRequestException{
		Code:        http.StatusBadRequest,
		Message:     http.StatusText(http.StatusBadRequest),
		Description: "The profile could not be parse.",
		ErrorData:   err,
	}
}

func (profileRequestException *ProfileRequestException) Error() string {
	return fmt.Sprintf("[CUSTOM],%d,%s?%s", profileRequestException.Code, static.MessageFormatWithDescription(enum.ResponseType("FAILURE").Value(), profileRequestException.Message, profileRequestException.Description), profileRequestException.ErrorData)
}

func (profileRequestException *ProfileRequestException) Throw() {
	exception.ThrowError(profileRequestException)
}
