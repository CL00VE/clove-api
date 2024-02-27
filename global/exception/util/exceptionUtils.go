// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package util

import "strings"

// tag extract util (by error message)
func TagExtract(str string) (string, string) {
	startIndex := strings.Index(str, "[")
	endIndex := strings.Index(str, "]")

	if startIndex == -1 || endIndex == -1 || startIndex+1 >= endIndex {
		return "SYSTEM", str
	}

	return str[startIndex+1 : endIndex], str[endIndex+2:]
}
