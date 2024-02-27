package exception

import (
	"clove-api/global/enum"
	"clove-api/global/exception"
	"clove-api/global/static"
	"fmt"
	"net/http"
)

type CommentNotFoundException struct {
	Code        int
	Message     string
	Description string
	ErrorData   string
}

func NewCommentNotFoundException(err string) *CommentNotFoundException {
	return &CommentNotFoundException{
		Code:        http.StatusNotFound,
		Message:     http.StatusText(http.StatusNotFound),
		Description: "The comment could not be found.",
		ErrorData:   err,
	}
}

func (commentNotFoundException *CommentNotFoundException) Error() string {
	return fmt.Sprintf("[CUSTOM],%d,%s?%s", commentNotFoundException.Code, static.MessageFormatWithDescription(enum.ResponseType("FAILURE").Value(), commentNotFoundException.Message, commentNotFoundException.Description), commentNotFoundException.ErrorData)
}

func (commentNotFoundException *CommentNotFoundException) Throw() {
	exception.ThrowError(commentNotFoundException)
}
