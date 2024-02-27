// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package status

import (
	"clove-api/global/enum"
	"clove-api/global/exception/check"
	"os"
	"strconv"
)

// A boolean variable that stores whether or not to use custom code.
var use bool

// init
func Setting() error {
	var err error

	code := os.Getenv("CLOVE_CODE")
	use, err = strconv.ParseBool(code)
	if err != nil {
		return err
	}

	return nil
}

// The GetCloveErrorCode function is a function that returns a custom error code.
func GetCloveErrorCode(httpCode int, errorType string) int {
	if use {
		return enum.ErrorType(errorType).Value() + httpCode
	}
	return httpCode
}

// The GetCloveSuccessCode function is a function that returns a custom success code.
func GetCloveSuccessCode(httpCode int) int {
	if use {
		customCode := os.Getenv("SUCCESS_CODE")
		CODE, strconvError := strconv.Atoi(customCode)
		check.SniffError(strconvError, strconvError)
		return CODE + httpCode
	}
	return httpCode
}
