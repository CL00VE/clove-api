package presentation

import (
	"clove-api/application/status"
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

func NewProfileController(profileService *service.ProfileService) *ProfileController {
	return &ProfileController{profileService: profileService}
}

type ProfileController struct {
	profileService *service.ProfileService
}

// Profile Create Controller
func (pc *ProfileController) Create(c *fiber.Ctx) error {
	requestBody, requestError := new(dto.ProfileCreateRequest).ParseX(c)
	check.SniffError(requestError, exception.NewProfileRequestException(requestError))

	_, serviceError := pc.profileService.Create(c.Context(), requestBody)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusCreated).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusCreated),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusCreated)),
	})
}

// Profile List Read Controller
func (pc *ProfileController) ReadAll(c *fiber.Ctx) error {
	profiles, serviceError := pc.profileService.ReadAll(c.Context())
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
		Data:    profiles,
	})
}

// Profile Read Controller
func (pc *ProfileController) Read(c *fiber.Ctx) error {
	id := c.Params("id")
	profile, serviceError := pc.profileService.Read(c.Context(), id)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
		Data:    profile,
	})
}

// Profile Update Controller
func (pc *ProfileController) Update(c *fiber.Ctx) error {
	requestBody, requestError := new(dto.ProfileUpdateRequest).ParseX(c)
	check.SniffError(requestError, exception.NewProfileRequestException(requestError))

	_, serviceError := pc.profileService.Update(c.Context(), requestBody)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusNoContent).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusNoContent),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusNoContent)),
	})
}
