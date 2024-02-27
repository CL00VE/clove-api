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

	"clove-api/domain/community/service"

	"github.com/gofiber/fiber/v2"
)

func NewPostController(postService *service.PostService) *PostController {
	return &PostController{postService: postService}
}

type PostController struct {
	postService *service.PostService
}

// Post Create Controller
func (pc *PostController) Create(c *fiber.Ctx) error {
	requestBody, requestError := new(dto.PostCreateRequest).ParseX(c)
	check.SniffError(requestError, exception.NewPostRequestException(requestError.Error()))

	_, serviceError := pc.postService.Create(c.Context(), requestBody)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusCreated).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusCreated),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusCreated)),
	})
}

// Post List Read Controller
func (pc *PostController) ReadAll(c *fiber.Ctx) error {
	order := c.Query("order")

	var posts []*ent.Post
	var serviceError error

	switch order {
	case "recent":
		posts, serviceError = pc.postService.ReadAllInRecentOrder(c.Context())
	case "popular":
		posts, serviceError = pc.postService.ReadAllInPopularOrder(c.Context())
	case "comments":
		posts, serviceError = pc.postService.ReadAllInCommentsOrder(c.Context())
	case "recommend":
		posts, serviceError = pc.postService.ReadAllInRandomOrder(c.Context())
	default:
		posts, serviceError = pc.postService.ReadAll(c.Context())
	}
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
		Data:    posts,
	})
}

// Post Read Controller
func (pc *PostController) Read(c *fiber.Ctx) error {
	id := c.Params("id")
	post, serviceError := pc.postService.Read(c.Context(), id)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
		Data:    post,
	})
}

// Post like add 1 Controller
func (pc *PostController) AddLike(c *fiber.Ctx) error {
	id := c.Params("id")
	_, serviceError := pc.postService.AddLike(c.Context(), id)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusNoContent).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusNoContent),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusNoContent)),
	})
}

// Post Update Controller
func (pc *PostController) Update(c *fiber.Ctx) error {
	requestBody, requestError := new(dto.PostUpdateRequest).ParseX(c)
	check.SniffError(requestError, exception.NewPostRequestException(requestError.Error()))

	_, serviceError := pc.postService.Update(c.Context(), requestBody)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusNoContent).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusNoContent),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusNoContent)),
	})
}

// Post List Delete Controller
func (pc *PostController) DeleteAll(c *fiber.Ctx) error {
	_, serviceError := pc.postService.DeleteAll(c.Context())
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusNoContent).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusNoContent),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusNoContent)),
	})
}

// Post Delete Controller
func (pc *PostController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	serviceError := pc.postService.Delete(c.Context(), id)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusNoContent).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusNoContent),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusNoContent)),
	})
}
