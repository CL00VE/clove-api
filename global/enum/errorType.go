// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package enum

import (
	"clove-api/global/exception/check"
	"os"
	"strconv"
)

// ErrorType Enum
type ErrorType string

// key
const (
	CUSTOM        ErrorType = "CUSTOM"
	SYSTEM        ErrorType = "SYSTEM"
	ERROR_HANDLER ErrorType = "ERROR_HANDLER"
)

// value
var errorTypeValues = map[ErrorType]int{}

// init
func ErrorTypeSetting() {
	CUSTOM_CODE := os.Getenv("CUSTOM_CODE")
	custom, strconvError := strconv.Atoi(CUSTOM_CODE)
	check.SniffError(strconvError, strconvError)

	SYSTEM_CODE := os.Getenv("SYSTEM_CODE")
	system, strconvError := strconv.Atoi(SYSTEM_CODE)
	check.SniffError(strconvError, strconvError)

	ERROR_HANDLER_CODE := os.Getenv("ERROR_HANDLER_CODE")
	errorHandler, strconvError := strconv.Atoi(ERROR_HANDLER_CODE)
	check.SniffError(strconvError, strconvError)

	errorTypeValues = map[ErrorType]int{
		CUSTOM:        custom,
		SYSTEM:        system,
		ERROR_HANDLER: errorHandler,
	}
}

// func
func (errorType ErrorType) Value() int {
	return errorTypeValues[errorType]
}
