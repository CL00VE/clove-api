package service

import (
	"clove-api/domain/community/domain/ent"
	"clove-api/domain/community/domain/repository"
	"clove-api/domain/community/exception"
	"clove-api/domain/community/presentation/dto"
	"clove-api/global/exception/check"
	"context"
	"strconv"
)

func NewProfileService(profileRepository *repository.ProfileRepository) *ProfileService {
	return &ProfileService{profileRepository: profileRepository}
}

type ProfileService struct {
	profileRepository *repository.ProfileRepository
}

// Profile Create Service
func (ps *ProfileService) Create(ctx context.Context, profileCreateRequest *dto.ProfileCreateRequest) (*ent.Profile, error) {
	profile, repositoryError := ps.profileRepository.Save(ctx, profileCreateRequest.ToEntity())
	return profile, repositoryError
}

// Profile List Read Service
func (ps *ProfileService) ReadAll(ctx context.Context) ([]*ent.Profile, error) {
	profiles, repositoryError := ps.profileRepository.FindAll(ctx)
	return profiles, repositoryError
}

// Profile Read Service
func (ps *ProfileService) Read(ctx context.Context, id string) (*ent.Profile, error) {
	ID, strconvError := strconv.Atoi(id)
	check.SniffError(strconvError, strconvError)
	profile, repositoryError := ps.profileRepository.FindById(ctx, ID)
	check.SniffNotFound(profile, exception.NewProfileNotFoundException(repositoryError))
	return profile, repositoryError
}

// Profile Update Service
func (ps *ProfileService) Update(ctx context.Context, profileUpdateRequest *dto.ProfileUpdateRequest) (*ent.Profile, error) {
	profile, repositoryError := ps.profileRepository.Save(ctx, profileUpdateRequest.ToEntity())
	return profile, repositoryError
}
