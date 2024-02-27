package exception

import (
	"clove-api/global/enum"
	"clove-api/global/exception"
	"clove-api/global/static"
	"fmt"
	"net/http"
)

type CommentRequestException struct {
	Code        int
	Message     string
	Description string
	ErrorData   string
}

func NewCommentRequestException(err string) *CommentRequestException {
	return &CommentRequestException{
		Code:        http.StatusBadRequest,
		Message:     http.StatusText(http.StatusBadRequest),
		Description: "The comment could not be parse.",
		ErrorData:   err,
	}
}

func (commentRequestException *CommentRequestException) Error() string {
	return fmt.Sprintf("[CUSTOM],%d,%s?%s", commentRequestException.Code, static.MessageFormatWithDescription(enum.ResponseType("FAILURE").Value(), commentRequestException.Message, commentRequestException.Description), commentRequestException.ErrorData)
}

func (commentRequestException *CommentRequestException) Throw() {
	exception.ThrowError(commentRequestException)
}
