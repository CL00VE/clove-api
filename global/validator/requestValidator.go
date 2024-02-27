// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func init() {
	v = validator.New()
}

// Request data validate util
func RequestValidator(request any) error {
	if err := v.Struct(request); err != nil {
		var messages []string
		for _, err := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("'%s': %s", err.Field(), err.Tag())
			messages = append(messages, message)
		}
		errMessage := strings.Join(messages, ", ")
		return fmt.Errorf("%s", errMessage)
	}

	return nil
}
