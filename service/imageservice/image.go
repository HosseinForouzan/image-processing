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
	Remove(ctx context.Context, fileName string) error
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

func (s Service) Upload(ctx context.Context, req param.UploadImageRequest) (param.UploadImageResponse, error) {
	imageFile, err := s.Storage.Save(ctx, req.Image)
	if err != nil {
		return param.UploadImageResponse{}, fmt.Errorf("can't save image")
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
		sErr := s.Storage.Remove(ctx, imageFile.OriginalKey)
		if sErr != nil {
			return param.UploadImageResponse{},fmt.Errorf("error in remove file: %w", sErr)
		}
		return param.UploadImageResponse{}, fmt.Errorf("unexpected error:%w", err)
	}



	return param.UploadImageResponse{
		ID: imageRepo.ID,
		OriginalName: imageFile.OriginalName,
		OriginalKey: imageFile.OriginalKey,
		ThumbnailKey: imageRepo.ThumbnailKey,
		ContentType: imageFile.ContentType,
		Size: imageFile.Size,
		Status: imageRepo.Status,

	}, nil



}