// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package handler

import (
	"clove-api/application/status"
	"clove-api/global/enum"
	"clove-api/global/exception/util"
	"clove-api/global/response"
	"clove-api/global/static"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// The ExceptionHandler function is a handler that handles exceptions or errors and custom errors that occur while the server is running.
func ExceptionHandler(c *fiber.Ctx, err error) error {
	tag, errMessage := util.TagExtract(err.Error())

	if tag == "CUSTOM" {
		index := strings.Index(errMessage, "?")
		data := errMessage[index+1:]
		parts := strings.Split(errMessage[:index], ",")
		code := http.StatusInternalServerError
		if str, err := strconv.Atoi(parts[0]); err == nil {
			code = str
		}
		return c.Status(code).JSON(&response.GeneralResponse{
			Status:  status.GetCloveErrorCode(code, tag),
			Message: parts[1],
			Data:    data,
		})
	}

	if tag == "SYSTEM" {
		code := http.StatusInternalServerError
		return c.Status(code).JSON(&response.GeneralResponse{
			Status:  status.GetCloveErrorCode(code, tag),
			Message: static.MessageFormat(enum.ResponseType("FAILURE").Value(), http.StatusText(code)),
			Data:    err.Error(),
		})
	}

	code := http.StatusInternalServerError
	return c.Status(code).JSON(&response.GeneralResponse{
		Status:  status.GetCloveErrorCode(code, "ERROR_HANDLER"),
		Message: static.MessageFormatWithDescription(enum.ResponseType("FAILURE").Value(), http.StatusText(code), "ExceptionHandler function operation error."),
		Data:    "An error occurred because the ExceptionHandler did not work properly.",
	})
}
