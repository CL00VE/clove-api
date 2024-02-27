// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package check

import "clove-api/global/exception"

// error processing
func SniffError(checkError error, throwError error) {
	if checkError != nil {
		exception.ThrowError(throwError)
	}
}

// not found processing
func SniffNotFound(checkData any, throwError error) {
	if checkData == nil {
		exception.ThrowError(throwError)
	}
}
