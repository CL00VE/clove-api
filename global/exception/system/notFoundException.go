// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package system

import (
	"clove-api/global/exception"
	"fmt"
	"net/http"
)

type NotFoundException struct {
	Code      int
	Message   string
	ErrorData string
}

func NewNotFoundException(url string) *NotFoundException {
	return &NotFoundException{
		Code:      http.StatusNotFound,
		Message:   http.StatusText(http.StatusNotFound),
		ErrorData: fmt.Sprintf("Address '%s' does not exist.", url),
	}
}

func (notFoundException *NotFoundException) Error() string {
	return notFoundException.ErrorData
}

func (notFoundException *NotFoundException) Throw() {
	exception.ThrowError(notFoundException)
}
