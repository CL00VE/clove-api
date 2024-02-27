package dto

import (
	"clove-api/domain/community/domain/ent"
	"clove-api/global/validator"

	"github.com/gofiber/fiber/v2"
)

type CommentCreateRequest struct {
	Text   string `json:"text" validate:"required"`
	PostID int    `json:"postID" validate:"required"`
}

func (request *CommentCreateRequest) ToEntity() *ent.Comment {
	comment := new(ent.Comment)
	comment.ID = 0
	comment.Text = request.Text
	comment.PostID = request.PostID
	return comment
}

func (request *CommentCreateRequest) ParseX(c *fiber.Ctx) (*CommentCreateRequest, error) {
	if err := validator.RequestValidator(request); err != nil {
		return nil, err
	}

	return request, nil
}

type CommentUpdateRequest struct {
	ID   int    `json:"ID" validate:"required"`
	Text string `json:"text"`
	Like int    `json:"like"`
}

func (request *CommentUpdateRequest) ToEntity() *ent.Comment {
	comment := new(ent.Comment)
	comment.ID = request.ID
	comment.Text = request.Text
	comment.Like = request.Like
	return comment
}

func (request *CommentUpdateRequest) ParseX(c *fiber.Ctx) (*CommentUpdateRequest, error) {
	if err := validator.RequestValidator(request); err != nil {
		return nil, err
	}

	return request, nil
}
