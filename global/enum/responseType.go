// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package enum

import "os"

// ResponseType Enum
type ResponseType string

// key
const (
	SUCCESS ResponseType = "SUCCESS"
	FAILURE ResponseType = "FAILURE"
)

// value
var responseTypeValues = map[ResponseType]string{
	SUCCESS: os.Getenv("SUCCESS"),
	FAILURE: os.Getenv("FAILURE"),
}

// func
func (responseType ResponseType) Value() string {
	return responseTypeValues[responseType]
}
