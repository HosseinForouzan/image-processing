package imageservice

import (
	"context"
	"fmt"
	"image_processing/entity"
	"image_processing/param"
	"mime/multipart"
)

type ImageRepository interface {
}

type Storage interface {
	Save(ctx context.Context ,fileHeader *multipart.FileHeader) (entity.Image, error)
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

func (s Service) SaveImage(ctx context.Context, req param.SaveImageRequest) (param.SaveImageResponse, error) {
	imageFile, err := s.Storage.Save(ctx, req.FileHeader)
	if err != nil {
		return param.SaveImageResponse{}, fmt.Errorf("can't save image")
	}

	return param.SaveImageResponse{
		OriginalName: imageFile.OriginalName,
		OriginalKey: imageFile.OriginalKey,
		ContentType: imageFile.ContentType,
		Size: imageFile.Size,

	}, nil



}