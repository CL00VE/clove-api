// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package presentation

import (
	"clove-api/application/status"
	"clove-api/domain/image/exception"
	"clove-api/domain/image/presentation/dto"
	"clove-api/domain/image/service"
	"clove-api/global/enum"
	"clove-api/global/exception/check"
	"clove-api/global/response"
	"clove-api/global/static"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func NewImageController(imageService *service.ImageService) *ImageController {
	return &ImageController{imageService: imageService}
}

type ImageController struct {
	imageService *service.ImageService
}

// Image Create Controller
func (ic *ImageController) Create(c *fiber.Ctx) error {
	log.Print("ImageController - Create")
	requestBody, requestError := new(dto.ImageCreateRequest).Parse(c)
	log.Print(requestBody, requestError)
	check.SniffError(requestError, exception.NewImageRequestException(requestError))
	log.Print("...")

	image, serviceError := ic.imageService.Create(c.Context(), requestBody)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusCreated).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusCreated),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusCreated)),
		Data:    image,
	})
}

// Image View Controller
func (ic *ImageController) GetInstanceByFileName(c *fiber.Ctx) error {
	fileName := c.Params("fileName")
	instance, imageFileName := ic.imageService.GetInstanceByFileName(c.Context(), fileName)
	c.Set("Content-Type", fmt.Sprintf("image/%s", filepath.Ext(imageFileName)[1:]))
	return c.Status(http.StatusOK).Send(instance)
}

// Image List Read Controller
func (ic *ImageController) ReadAll(c *fiber.Ctx) error {
	images, serviceError := ic.imageService.ReadAll(c.Context())
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
		Data:    images,
	})
}

// Image Read Controller
func (ic *ImageController) Read(c *fiber.Ctx) error {
	id := c.Params("id")
	image, serviceError := ic.imageService.Read(c.Context(), id)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
		Data:    image,
	})
}

// Image Read Controller (condition: name)
func (ic *ImageController) ReadByName(c *fiber.Ctx) error {
	name := c.Params("name")
	image, serviceError := ic.imageService.ReadByName(c.Context(), name)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
		Data:    image,
	})
}

// Image Read Controller (condition: fileName)
func (ic *ImageController) ReadByFileName(c *fiber.Ctx) error {
	fileName := c.Params("fileName")
	image, serviceError := ic.imageService.ReadByFileName(c.Context(), fileName)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
		Data:    image,
	})
}

// Image List Delete Controller
func (ic *ImageController) DeleteAll(c *fiber.Ctx) error {
	_, serviceError := ic.imageService.DeleteAll(c.Context())
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
	})
}

// Image Delete Controller
func (ic *ImageController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	serviceError := ic.imageService.Delete(c.Context(), id)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
	})
}

// Image Delete Controller (condition: name)
func (ic *ImageController) DeleteByName(c *fiber.Ctx) error {
	name := c.Params("name")
	serviceError := ic.imageService.DeleteByName(c.Context(), name)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
	})
}

// Image Delete Controller (condition: fileName)
func (ic *ImageController) DeleteByFileName(c *fiber.Ctx) error {
	fileName := c.Params("fileName")
	serviceError := ic.imageService.DeleteByFileName(c.Context(), fileName)
	check.SniffError(serviceError, serviceError)

	return c.Status(http.StatusOK).JSON(&response.GeneralResponse{
		Status:  status.GetCloveSuccessCode(http.StatusOK),
		Message: static.MessageFormat(enum.ResponseType("SUCCESS").Value(), http.StatusText(http.StatusOK)),
	})
}
