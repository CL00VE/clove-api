package dto

import (
	"clove-api/domain/community/domain/ent"
	"clove-api/global/validator"

	"github.com/gofiber/fiber/v2"
)

type ProfileCreateRequest struct {
	Name  string `form:"name" validate:"required"`
	Image string `json:"image" validate:"required"`
}

func (request *ProfileCreateRequest) ToEntity() *ent.Profile {
	profile := new(ent.Profile)
	profile.ID = 0
	profile.Name = request.Name
	profile.Image = request.Image
	return profile
}

func (request *ProfileCreateRequest) ParseX(c *fiber.Ctx) (*ProfileCreateRequest, error) {
	if err := validator.RequestValidator(request); err != nil {
		return nil, err
	}

	return request, nil
}

type ProfileUpdateRequest struct {
	ID    int    `json:"ID" validate:"required"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (request *ProfileUpdateRequest) ToEntity() *ent.Profile {
	profile := new(ent.Profile)
	profile.Name = request.Name
	profile.Image = request.Image
	return profile
}

func (request *ProfileUpdateRequest) ParseX(c *fiber.Ctx) (*ProfileUpdateRequest, error) {
	if err := validator.RequestValidator(request); err != nil {
		return nil, err
	}

	return request, nil
}
