package dto

import (
	"clove-api/domain/image/domain/ent"
	"clove-api/global/util"
	"clove-api/global/validator"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ImageCreateRequest struct {
	Name        string                `form:"name" validate:"required"`
	Description string                `form:"description"`
	Image       *multipart.FileHeader `form:"image" validate:"image,required"`
}

func (request *ImageCreateRequest) ToEntity() (*ent.Image, error) {
	image := new(ent.Image)
	image.ID = 0
	image.Name = request.Name
	image.Description = request.Description
	image.FileName = fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(request.Image.Filename))
	image.URL = fmt.Sprintf("%s://%s:%s/%s/api/image/instance/%s", os.Getenv("SERVER_PROTOCOL"), os.Getenv("SERVER_DOMAIN"), os.Getenv("SERVER_PORT"), os.Getenv("SERVER_VERSION"), image.FileName)
	image.Size = util.FormatFileSize(request.Image.Size)
	fileBytes, err := util.GetByteByFile(request.Image)
	if err != nil {
		return nil, err
	}
	image.Instance = fileBytes
	return image, nil
}

func (request *ImageCreateRequest) Parse(c *fiber.Ctx) (*ImageCreateRequest, error) {
	request.Name = c.FormValue("name")
	request.Description = c.FormValue("description")

	var err error
	if request.Image, err = c.FormFile("image"); err != nil {
		return nil, err
	}

	if err := validator.RequestValidator(request); err != nil {
		return nil, err
	}

	return request, nil
}
