// 🩷 CLOVE provides services to check and manage one's psychology or mood through psychological tests and to engage in social activities through the community.
// 🛠️ Github Repository: https://github.com/CL00VE/clove-api
// 🗒️ Project Portfolio: https://chill-brisket-c2b.notion.site/CLOVE-661adfae4e9b48f69a0b55dae0af17d9
package repository

import (
	"clove-api/domain/image/domain/ent"
	"clove-api/domain/image/domain/ent/image"
	"clove-api/domain/image/exception"
	"clove-api/global/exception/check"
	"context"
)

func NewImageRepository(db *ent.Client) *ImageRepository {
	return &ImageRepository{
		db: db,
	}
}

type ImageRepository struct {
	db *ent.Client
}

func (ir *ImageRepository) Save(ctx context.Context, image *ent.Image) (*ent.Image, error) {
	return ir.db.Image.Create().
		SetName(image.Name).
		SetDescription(image.Description).
		SetFileName(image.FileName).
		SetURL(image.URL).
		SetSize(image.Size).
		SetInstance(image.Instance).
		Save(ctx)
}

func (ir *ImageRepository) FindAll(ctx context.Context) ([]*ent.Image, error) {
	return ir.db.Image.Query().All(ctx)
}

func (ir *ImageRepository) FindById(ctx context.Context, id int) (*ent.Image, error) {
	return ir.db.Image.Query().Where(image.ID(id)).Only(ctx)
}

func (ir *ImageRepository) FindByName(ctx context.Context, name string) (*ent.Image, error) {
	return ir.db.Image.Query().Where(image.Name(name)).Only(ctx)
}

func (ir *ImageRepository) FindByFileName(ctx context.Context, fileName string) (*ent.Image, error) {
	return ir.db.Image.Query().Where(image.FileName(fileName)).Only(ctx)
}

func (ir *ImageRepository) RemoveAll(ctx context.Context) (int, error) {
	return ir.db.Image.Delete().Exec(ctx)
}

func (ir *ImageRepository) RemoveById(ctx context.Context, id int) error {
	return ir.db.Image.DeleteOneID(id).Exec(ctx)
}

func (ir *ImageRepository) RemoveByName(ctx context.Context, name string) error {
	image, dbError := ir.db.Image.Query().Where(image.Name(name)).Only(ctx)
	check.SniffNotFound(image, exception.NewImageNotFoundException(dbError.Error()))
	check.SniffError(dbError, dbError)
	return ir.db.Image.DeleteOne(image).Exec(ctx)
}

func (ir *ImageRepository) RemoveByFileName(ctx context.Context, fileName string) error {
	image, dbError := ir.db.Image.Query().Where(image.FileName(fileName)).Only(ctx)
	check.SniffNotFound(image, exception.NewImageNotFoundException(dbError.Error()))
	check.SniffError(dbError, dbError)
	return ir.db.Image.DeleteOne(image).Exec(ctx)
}
