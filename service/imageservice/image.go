package imageservice

import (
	"context"
	"fmt"
	"image_processing/constant"
	"image_processing/entity"
	"image_processing/param"
	"mime/multipart"
)

type ImageRepository interface {
	Save(ctx context.Context, image entity.Image) (entity.Image, error)
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


	imageEntity := entity.Image{
		OriginalName: imageFile.OriginalName,
		OriginalKey: imageFile.OriginalKey,
		ThumbnailKey: "",
		ContentType: imageFile.ContentType,
		Size: imageFile.Size,
		Status: constant.PROCESSING,

	}

	imageRepo, err := s.ImageRepo.Save(ctx, imageEntity)
	if err != nil {
		return param.SaveImageResponse{}, fmt.Errorf("unexpected error:%w", err)
	}



	return param.SaveImageResponse{
		ID: imageRepo.ID,
		OriginalName: imageFile.OriginalName,
		OriginalKey: imageFile.OriginalKey,
		ThumbnailKey: imageRepo.ThumbnailKey,
		ContentType: imageFile.ContentType,
		Size: imageFile.Size,
		Status: imageRepo.Status,

	}, nil



}