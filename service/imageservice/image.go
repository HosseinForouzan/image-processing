package imageservice

import (
	"image_processing/entity"
	"mime/multipart"
)

type ImageRepository interface {
}

type Storage interface {
	Save(fileHeader *multipart.FileHeader) (entity.Image, error)
}

type Service struct {
	ImageRepo ImageRepository
	Storage   Storage
}

func New(imageRepo ImageRepository, storage Storage) Service {
	return Service{
		ImageRepo: imageRepo,
		Storage:   storage,
	}
}
