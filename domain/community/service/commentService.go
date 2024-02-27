package service

import (
	"clove-api/domain/community/domain/ent"
	"clove-api/domain/community/domain/repository"
	"clove-api/domain/community/exception"
	"clove-api/domain/community/presentation/dto"
	"clove-api/global/exception/check"
	"context"
	"sort"
	"strconv"
)

func NewCommentService(commentRepository *repository.CommentRepository) *CommentService {
	return &CommentService{commentRepository: commentRepository}
}

type CommentService struct {
	commentRepository *repository.CommentRepository
}

// Comment Create Service
func (cs *CommentService) Create(ctx context.Context, commentCreateRequest *dto.CommentCreateRequest) (*ent.Comment, error) {
	comment, repositoryError := cs.commentRepository.Save(ctx, commentCreateRequest.ToEntity())
	return comment, repositoryError
}

// Comment List Read Service
func (cs *CommentService) ReadAll(ctx context.Context) ([]*ent.Comment, error) {
	comments, repositoryError := cs.commentRepository.FindAll(ctx)
	return comments, repositoryError
}

// Comment List Read Service
func (cs *CommentService) ReadAllInRecentOrder(ctx context.Context) ([]*ent.Comment, error) {
	comments, repositoryError := cs.commentRepository.FindAll(ctx)
	if repositoryError != nil {
		return nil, repositoryError
	}

	sort.Slice(comments, func(i, j int) bool {
		return comments[i].CreatedAt.After(comments[j].CreatedAt)
	})
	return comments, nil
}

// Comment List Read Service
func (cs *CommentService) ReadAllByPostID(ctx context.Context, postID int) ([]*ent.Comment, error) {
	comments, repositoryError := cs.commentRepository.FindAllByPostID(ctx, postID)
	return comments, repositoryError
}

// Comment List Read Service
func (cs *CommentService) ReadAllByPostIDInRecentOrder(ctx context.Context, postID int) ([]*ent.Comment, error) {
	comments, repositoryError := cs.commentRepository.FindAllByPostID(ctx, postID)
	if repositoryError != nil {
		return nil, repositoryError
	}

	sort.Slice(comments, func(i, j int) bool {
		return comments[i].CreatedAt.After(comments[j].CreatedAt)
	})
	return comments, nil
}

// Comment Read Service
func (cs *CommentService) Read(ctx context.Context, id string) (*ent.Comment, error) {
	ID, strconvError := strconv.Atoi(id)
	check.SniffError(strconvError, strconvError)
	comment, repositoryError := cs.commentRepository.FindById(ctx, ID)
	check.SniffNotFound(comment, exception.NewCommentNotFoundException(repositoryError.Error()))
	return comment, repositoryError
}

// Comment like add 1 Service
func (cs *CommentService) AddLike(ctx context.Context, id string) (*ent.Comment, error) {
	ID, strconvError := strconv.Atoi(id)
	check.SniffError(strconvError, strconvError)
	comment, repositoryError := cs.commentRepository.AddLikeById(ctx, ID)
	return comment, repositoryError
}

// Comment Update Service
func (cs *CommentService) Update(ctx context.Context, CommentUpdateRequest *dto.CommentUpdateRequest) (*ent.Comment, error) {
	comment, repositoryError := cs.commentRepository.Save(ctx, CommentUpdateRequest.ToEntity())
	return comment, repositoryError
}

// Comment List Delete Service
func (cs *CommentService) DeleteAll(ctx context.Context) (int, error) {
	return cs.commentRepository.RemoveAll(ctx)
}

// Comment Delete Service
func (cs *CommentService) Delete(ctx context.Context, id string) error {
	ID, strconvError := strconv.Atoi(id)
	check.SniffError(strconvError, strconvError)
	return cs.commentRepository.RemoveById(ctx, ID)
}
