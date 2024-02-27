package repository

import (
	"clove-api/domain/community/domain/ent"
	"clove-api/domain/community/domain/ent/comment"
	"clove-api/global/util"
	"context"
	"math/rand"
	"time"
)

func NewCommentRepository(db *ent.Client) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

type CommentRepository struct {
	db *ent.Client
}

func (cr *CommentRepository) Save(ctx context.Context, comment *ent.Comment) (*ent.Comment, error) {
	if comment.ID == 0 {
		rand.Seed(time.Now().UnixNano())

		profileCount, err := cr.db.Profile.Query().Count(ctx)
		if err != nil {
			return nil, err
		}

		randomOffset := rand.Intn(profileCount)

		randomProfile, err := cr.db.Profile.Query().
			Offset(randomOffset).
			Limit(1).
			Only(ctx)
		if err != nil {
			return nil, err
		}

		// Comment doesn't have an ID, so it's a new post. Create it.
		return cr.db.Comment.Create().
			SetText(comment.Text).
			SetPostID(comment.PostID).
			SetProfile(randomProfile).
			Save(ctx)
	}

	findedComment, err := cr.db.Comment.Get(ctx, comment.ID)
	if err != nil {
		return nil, err
	}

	// Comment has an ID, so it already exists. Update it.
	return cr.db.Comment.UpdateOneID(comment.ID).
		SetText(util.Ternary(comment.Text == "", findedComment.Text, comment.Text).(string)).
		SetLike(util.Ternary(comment.Like == -1, findedComment.Like, comment.Like).(int)).
		Save(ctx)
}

func (cr *CommentRepository) FindAll(ctx context.Context) ([]*ent.Comment, error) {
	return cr.db.Comment.Query().WithProfile().All(ctx)
}

func (cr *CommentRepository) FindAllByPostID(ctx context.Context, postID int) ([]*ent.Comment, error) {
	return cr.db.Comment.Query().Where(comment.PostIDEQ(postID)).WithProfile().All(ctx)
}

func (cr *CommentRepository) FindById(ctx context.Context, id int) (*ent.Comment, error) {
	return cr.db.Comment.Query().Where(comment.ID(id)).WithProfile().Only(ctx)
}

func (cr *CommentRepository) AddLikeById(ctx context.Context, id int) (*ent.Comment, error) {
	return cr.db.Comment.UpdateOneID(id).AddLike(1).Save(ctx)
}

func (cr *CommentRepository) RemoveAll(ctx context.Context) (int, error) {
	return cr.db.Comment.Delete().Exec(ctx)
}

func (cr *CommentRepository) RemoveById(ctx context.Context, id int) error {
	return cr.db.Comment.DeleteOneID(id).Exec(ctx)
}
