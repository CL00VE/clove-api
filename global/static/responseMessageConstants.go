// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package static

import "fmt"

// Response Fotmat (with Description)
func MessageFormatWithDescription(responseType string, httpStatusText string, description string) string {
	return fmt.Sprintf("%s %s: %s", responseType, httpStatusText, description)
}

// Response Fotmat
func MessageFormat(responseType string, httpStatusText string) string {
	return fmt.Sprintf("%s %s", responseType, httpStatusText)
}
