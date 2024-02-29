package presentation

import (
	"clove-api/application/status"
	"clove-api/domain/community/domain/ent"
	"clove-api/domain/community/exception"
	"clove-api/domain/community/presentation/dto"
	"clove-api/global/enum"
	"clove-api/global/exception/check"
	"clove-api/global/response"
	"clove-api/global/static"
	"net/http"
	"strconv"

	"clove-api/domain/community/service"

	"github.com/gofiber/fiber/v2"
)

func NewCommentController(commentService *service.CommentService) *CommentController {
	return &CommentController{commentService: commentService}
}

type CommentController struct {
	commentService *service.CommentService
}

// Comment Create Controller
func (cc *CommentController) Create(c *fiber.Ctx) error {
	requestBody, requestError := new(dto.CommentCreateRequest).Parse(c)
	check.SniffError(requestError, exception.NewCommentRequestException(requestError))

	_, serviceError := cc.commentService.Create(c.Context(), requestBody)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusCreated).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusCreated),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusCreated)),
	})
}

// Comment List Read Controller
func (cc *CommentController) ReadAll(c *fiber.Ctx) error {
	postID := c.Query("postID")
	order := c.Query("order")

	var ID int
	var comments []*ent.Comment
	var serviceError error
	var strconvError error

	if postID == "" {
		switch order {
		case "recent":
			comments, serviceError = cc.commentService.ReadAllInRecentOrder(c.Context())
		default:
			comments, serviceError = cc.commentService.ReadAll(c.Context())
		}
	} else {
		ID, strconvError = strconv.Atoi(postID)
		check.SniffError(strconvError, strconvError)

		switch order {
		case "recent":
			comments, serviceError = cc.commentService.ReadAllByPostIDInRecentOrder(c.Context(), ID)
		default:
			comments, serviceError = cc.commentService.ReadAllByPostID(c.Context(), ID)
		}
	}
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
		Data:    comments,
	})
}

// Comment Read Controller
func (cc *CommentController) Read(c *fiber.Ctx) error {
	id := c.Params("id")
	comment, serviceError := cc.commentService.Read(c.Context(), id)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
		Data:    comment,
	})
}

// Comment like add 1 Controller
func (cc *CommentController) AddLike(c *fiber.Ctx) error {
	id := c.Params("id")
	_, serviceError := cc.commentService.AddLike(c.Context(), id)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
	})
}

// Comment Update Controller
func (cc *CommentController) Update(c *fiber.Ctx) error {
	requestBody, requestError := new(dto.CommentUpdateRequest).Parse(c)
	check.SniffError(requestError, exception.NewCommentRequestException(requestError))

	_, serviceError := cc.commentService.Update(c.Context(), requestBody)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
	})
}

// Comment List Delete Controller
func (cc *CommentController) DeleteAll(c *fiber.Ctx) error {
	_, serviceError := cc.commentService.DeleteAll(c.Context())
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
	})
}

// Comment Delete Controller
func (cc *CommentController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	serviceError := cc.commentService.Delete(c.Context(), id)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
	})
}
