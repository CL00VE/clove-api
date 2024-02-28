// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package global

import (
	"clove-api/global/enum"
	"clove-api/global/exception"
	"clove-api/global/static"
	"fmt"
	"net/http"
)

type ExternalAPIException struct {
	Code        int
	Message     string
	Description string
	ErrorData   string
}

func NewExternalAPIException(err error) *ExternalAPIException {
	data := ""
	if err != nil {
		data = err.Error()
	}
	return &ExternalAPIException{
		Code:        http.StatusBadGateway,
		Message:     http.StatusText(http.StatusBadGateway),
		Description: "External API server error.",
		ErrorData:   data,
	}
}

func (externalAPIException *ExternalAPIException) Error() string {
	return fmt.Sprintf("[CUSTOM],%d,%s?%s", externalAPIException.Code, static.MessageFormatWithDescription(enum.ResponseType("FAILURE").Value(), externalAPIException.Message, externalAPIException.Description), externalAPIException.ErrorData)
}

func (externalAPIException *ExternalAPIException) Throw() {
	exception.ThrowError(externalAPIException)
}
