package service

import (
	"clove-api/domain/community/domain/ent"
	"clove-api/domain/community/domain/repository"
	"clove-api/domain/community/exception"
	"clove-api/domain/community/presentation/dto"
	"clove-api/global/exception/check"
	"context"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func NewPostService(postRepository *repository.PostRepository) *PostService {
	return &PostService{postRepository: postRepository}
}

type PostService struct {
	postRepository *repository.PostRepository
}

// Post Create Service
func (ps *PostService) Create(ctx context.Context, postCreateRequest *dto.PostCreateRequest) (*ent.Post, error) {
	post, repositoryError := ps.postRepository.Save(ctx, postCreateRequest.ToEntity())
	return post, repositoryError
}

// Post List Read Service
func (ps *PostService) ReadAll(ctx context.Context) ([]*ent.Post, error) {
	posts, repositoryError := ps.postRepository.FindAll(ctx)
	return posts, repositoryError
}

// Post List Read Service
func (ps *PostService) ReadAllInRecentOrder(ctx context.Context) ([]*ent.Post, error) {
	posts, repositoryError := ps.postRepository.FindAll(ctx)
	if repositoryError != nil {
		return nil, repositoryError
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})

	return posts, nil
}

// Post List Read Service
func (ps *PostService) ReadAllInPopularOrder(ctx context.Context) ([]*ent.Post, error) {
	posts, repositoryError := ps.postRepository.FindAll(ctx)
	if repositoryError != nil {
		return nil, repositoryError
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Like > posts[j].Like
	})

	return posts, nil
}

// Post List Read Service
func (ps *PostService) ReadAllInCommentsOrder(ctx context.Context) ([]*ent.Post, error) {
	posts, repositoryError := ps.postRepository.FindAllWithComments(ctx)
	if repositoryError != nil {
		return nil, repositoryError
	}

	sort.Slice(posts, func(i, j int) bool {
		return len(posts[i].Edges.Comments) > len(posts[j].Edges.Comments)
	})

	return posts, nil
}

// Post List Read Service
func (ps *PostService) ReadAllInRandomOrder(ctx context.Context) ([]*ent.Post, error) {
	posts, repositoryError := ps.postRepository.FindAll(ctx)
	if repositoryError != nil {
		return nil, repositoryError
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(posts), func(i, j int) {
		posts[i], posts[j] = posts[j], posts[i]
	})

	return posts, nil
}

// Post Read Service
func (ps *PostService) Read(ctx context.Context, id string) (*ent.Post, error) {
	ID, strconvError := strconv.Atoi(id)
	check.SniffError(strconvError, strconvError)
	post, repositoryError := ps.postRepository.FindById(ctx, ID)
	check.SniffNotFound(post, exception.NewPostNotFoundException(repositoryError))
	return post, repositoryError
}

// Post like add 1 Service
func (ps *PostService) AddLike(ctx context.Context, id string) (*ent.Post, error) {
	ID, strconvError := strconv.Atoi(id)
	check.SniffError(strconvError, strconvError)
	post, repositoryError := ps.postRepository.AddLikeById(ctx, ID)
	return post, repositoryError
}

// Post Update Service
func (ps *PostService) Update(ctx context.Context, postUpdateRequest *dto.PostUpdateRequest) (*ent.Post, error) {
	post, repositoryError := ps.postRepository.Save(ctx, postUpdateRequest.ToEntity())
	return post, repositoryError
}

// Post List Delete Service
func (ps *PostService) DeleteAll(ctx context.Context) (int, error) {
	return ps.postRepository.RemoveAll(ctx)
}

// Post Delete Service
func (ps *PostService) Delete(ctx context.Context, id string) error {
	ID, strconvError := strconv.Atoi(id)
	check.SniffError(strconvError, strconvError)
	return ps.postRepository.RemoveById(ctx, ID)
}
