package dto

import (
	"clove-api/domain/community/domain/ent"
	"clove-api/global/validator"

	"github.com/gofiber/fiber/v2"
)

type PostCreateRequest struct {
	Icon    string `json:"icon" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (request *PostCreateRequest) ToEntity() *ent.Post {
	post := new(ent.Post)
	post.ID = 0
	post.Icon = request.Icon
	post.Content = request.Content
	return post
}

func (request *PostCreateRequest) ParseX(c *fiber.Ctx) (*PostCreateRequest, error) {
	if err := validator.RequestValidator(request); err != nil {
		return nil, err
	}

	return request, nil
}

type PostUpdateRequest struct {
	ID      int    `json:"ID" validate:"required"`
	Icon    string `json:"icon"`
	Content string `json:"content"`
	Like    int    `json:"like"`
}

func (request *PostUpdateRequest) ToEntity() *ent.Post {
	post := new(ent.Post)
	post.ID = request.ID
	post.Icon = request.Icon
	post.Content = request.Content
	post.Like = request.Like
	return post
}

func (request *PostUpdateRequest) ParseX(c *fiber.Ctx) (*PostUpdateRequest, error) {
	if err := validator.RequestValidator(request); err != nil {
		return nil, err
	}

	return request, nil
}
