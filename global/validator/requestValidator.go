// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package validator

import (
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/go-playground/validator/v10"
)

// var v *validator.Validate

// func init() {
// 	v = validator.New()
// }

var v *validator.Validate

var allowExtentions = []string{".jpeg", ".jpg", ".png", ".mov", ".svg"}

func init() {
	v = validator.New()

	v.RegisterValidation("image", func(fl validator.FieldLevel) bool {
		field := fl.Field()

		if fileHeader, ok := field.Interface().(multipart.FileHeader); ok {
			ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
			for _, allowedExt := range allowExtentions {
				if ext == allowedExt {
					return true
				}
			}
		}
		return false
	})
}

// Request data validate util
func RequestValidator(request any) error {
	log.Print("RequestValidator")
	if err := v.Struct(request); err != nil {
		var messages []string
		for _, err := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("'%s': %s", err.Field(), err.Tag())
			messages = append(messages, message)
		}
		errMessage := strings.Join(messages, ", ")
		return fmt.Errorf("%s", errMessage)
	}
	log.Print("success")

	return nil
}
