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

type TestException struct {
	Code        int
	Message     string
	Description string
	ErrorData   string
}

func NewTestException(err string) *TestException {
	return &TestException{
		Code:        http.StatusInternalServerError,
		Message:     http.StatusText(http.StatusInternalServerError),
		Description: "Error handler test.",
		ErrorData:   err,
	}
}

func (testException *TestException) Error() string {
	return fmt.Sprintf("[CUSTOM],%d,%s?%s", testException.Code, static.MessageFormatWithDescription(enum.ResponseType("FAILURE").Value(), testException.Message, testException.Description), testException.ErrorData)
}

func (testException *TestException) Throw() {
	exception.ThrowError(testException)
}
