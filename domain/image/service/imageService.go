// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package service

import (
	"clove-api/domain/image/domain/ent"
	"clove-api/domain/image/domain/repository"
	"clove-api/domain/image/exception"
	"clove-api/domain/image/presentation/dto"
	"clove-api/global/exception/check"
	"context"
	"strconv"
)

func NewImageService(imageRepository *repository.ImageRepository) *ImageService {
	return &ImageService{imageRepository: imageRepository}
}

type ImageService struct {
	imageRepository *repository.ImageRepository
}

// Image Create Service
func (is *ImageService) Create(ctx context.Context, imageCreateRequest *dto.ImageCreateRequest) (*ent.Image, error) {
	imageEntity, mappingError := imageCreateRequest.ToEntity()
	check.SniffError(mappingError, mappingError)

	image, repositoryError := is.imageRepository.Save(ctx, imageEntity)
	return image, repositoryError
}

// Image View Service
func (is *ImageService) GetInstanceByFileName(ctx context.Context, fileName string) ([]byte, error) {
	image, repositoryError := is.imageRepository.FindWithInstanceByFileName(ctx, fileName)
	check.SniffNotFound(image, exception.NewImageNotFoundException(repositoryError))

	return image.Instance, nil
}

// Image List Read Service
func (is *ImageService) ReadAll(ctx context.Context) ([]*ent.Image, error) {
	images, repositoryError := is.imageRepository.FindAll(ctx)
	return images, repositoryError
}

// Image Read Service
func (is *ImageService) Read(ctx context.Context, id string) (*ent.Image, error) {
	ID, strconvError := strconv.Atoi(id)
	check.SniffError(strconvError, strconvError)
	image, repositoryError := is.imageRepository.FindById(ctx, ID)
	check.SniffNotFound(image, exception.NewImageNotFoundException(repositoryError))
	return image, repositoryError
}

// Image Read Service (condition: name)
func (is *ImageService) ReadByName(ctx context.Context, name string) (*ent.Image, error) {
	image, repositoryError := is.imageRepository.FindByName(ctx, name)
	check.SniffNotFound(image, exception.NewImageNotFoundException(repositoryError))
	return image, repositoryError
}

// Image Read Service (condition: fileName)
func (is *ImageService) ReadByFileName(ctx context.Context, fileName string) (*ent.Image, error) {
	image, repositoryError := is.imageRepository.FindByFileName(ctx, fileName)
	check.SniffNotFound(image, exception.NewImageNotFoundException(repositoryError))
	return image, repositoryError
}

// Image List Delete Service
func (is *ImageService) DeleteAll(ctx context.Context) (int, error) {
	return is.imageRepository.RemoveAll(ctx)
}

// Image Delete Service
func (is *ImageService) Delete(ctx context.Context, id string) error {
	ID, strconvError := strconv.Atoi(id)
	check.SniffError(strconvError, strconvError)
	return is.imageRepository.RemoveById(ctx, ID)
}

// Image Delete Service (condition: name)
func (is *ImageService) DeleteByName(ctx context.Context, name string) error {
	return is.imageRepository.RemoveByName(ctx, name)
}

// Image Delete Service (condition: fileName)
func (is *ImageService) DeleteByFileName(ctx context.Context, fileName string) error {
	return is.imageRepository.RemoveByFileName(ctx, fileName)
}
