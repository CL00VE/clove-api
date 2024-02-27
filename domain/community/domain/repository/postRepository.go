package repository

import (
	"clove-api/domain/community/domain/ent"
	"clove-api/domain/community/domain/ent/comment"
	"clove-api/domain/community/domain/ent/post"
	"clove-api/global/util"
	"context"
	"math/rand"
	"time"
)

func NewPostRepository(db *ent.Client) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

type PostRepository struct {
	db *ent.Client
}

func (pr *PostRepository) Save(ctx context.Context, post *ent.Post) (*ent.Post, error) {
	if post.ID == 0 {
		rand.Seed(time.Now().UnixNano())

		profileCount, err := pr.db.Profile.Query().Count(ctx)
		if err != nil {
			return nil, err
		}

		randomOffset := rand.Intn(profileCount)

		randomProfile, err := pr.db.Profile.Query().
			Offset(randomOffset).
			Limit(1).
			Only(ctx)
		if err != nil {
			return nil, err
		}

		// Post doesn't have an ID, so it's a new post. Create it.
		return pr.db.Post.Create().
			SetIcon(post.Icon).
			SetContent(post.Content).
			SetProfile(randomProfile).
			Save(ctx)
	}

	findedPost, err := pr.db.Post.Get(ctx, post.ID)
	if err != nil {
		return nil, err
	}

	// Post has an ID, so it already exists. Update it.
	return pr.db.Post.UpdateOneID(post.ID).
		SetIcon(util.Ternary(post.Icon == "", findedPost.Icon, post.Icon).(string)).
		SetContent(util.Ternary(post.Content == "", findedPost.Content, post.Content).(string)).
		SetLike(util.Ternary(post.Like == -1, findedPost.Like, post.Like).(int)).
		Save(ctx)
}

func (pr *PostRepository) FindAll(ctx context.Context) ([]*ent.Post, error) {
	return pr.db.Post.Query().WithProfile().All(ctx)
}

func (pr *PostRepository) FindAllWithComments(ctx context.Context) ([]*ent.Post, error) {
	return pr.db.Post.Query().WithComments().All(ctx)
}

func (pr *PostRepository) FindById(ctx context.Context, id int) (*ent.Post, error) {
	return pr.db.Post.Query().Where(post.ID(id)).WithProfile().Only(ctx)
}

func (pr *PostRepository) AddLikeById(ctx context.Context, id int) (*ent.Post, error) {
	return pr.db.Post.UpdateOneID(id).AddLike(1).Save(ctx)
}

func (pr *PostRepository) RemoveAll(ctx context.Context) (int, error) {
	posts, err := pr.db.Post.Query().All(ctx)
	if err != nil {
		return 0, err
	}

	for _, post := range posts {
		if _, err := pr.db.Comment.Delete().Where(comment.PostIDEQ(post.ID)).Exec(ctx); err != nil {
			return 0, err
		}
	}

	return pr.db.Post.Delete().Exec(ctx)
}

func (pr *PostRepository) RemoveById(ctx context.Context, id int) error {
	if _, err := pr.db.Comment.Delete().Where(comment.PostIDEQ(id)).Exec(ctx); err != nil {
		return err
	}

	return pr.db.Post.DeleteOneID(id).Exec(ctx)
}
