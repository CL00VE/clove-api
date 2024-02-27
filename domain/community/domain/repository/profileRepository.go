package repository

import (
	"clove-api/domain/community/domain/ent"
	"clove-api/domain/community/domain/ent/profile"
	"clove-api/global/util"
	"context"
)

func NewProfileRepository(db *ent.Client) *ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

type ProfileRepository struct {
	db *ent.Client
}

func (pr *ProfileRepository) Save(ctx context.Context, profile *ent.Profile) (*ent.Profile, error) {
	if profile.ID == 0 {
		// Post doesn't have an ID, so it's a new post. Create it.
		return pr.db.Profile.Create().
			SetName(profile.Name).
			SetImage(profile.Image).
			Save(ctx)
	}

	findedProfile, err := pr.db.Profile.Get(ctx, profile.ID)
	if err != nil {
		return nil, err
	}

	// Post has an ID, so it already exists. Update it.
	return pr.db.Profile.UpdateOneID(profile.ID).
		SetName(util.Ternary(profile.Name == "", findedProfile.Name, profile.Name).(string)).
		SetImage(util.Ternary(profile.Image == "", findedProfile.Image, profile.Image).(string)).
		Save(ctx)
}

func (pr *ProfileRepository) FindAll(ctx context.Context) ([]*ent.Profile, error) {
	return pr.db.Profile.Query().All(ctx)
}

func (pr *ProfileRepository) FindById(ctx context.Context, id int) (*ent.Profile, error) {
	return pr.db.Profile.Query().Where(profile.ID(id)).Only(ctx)
}
