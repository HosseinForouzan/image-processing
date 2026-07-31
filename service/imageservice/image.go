package imageservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image_processing/constant"
	"image_processing/entity"
	"image_processing/event"
	"image_processing/param"
	"io"
	"path/filepath"


	"github.com/disintegration/imaging"
	"github.com/google/uuid"
)

type ImageRepository interface {
	Save(ctx context.Context, image entity.Image) (entity.Image, error)
	GetByID(ctx context.Context, id uint) (entity.Image, error)
	Update(ctx context.Context, image entity.Image) error
	Remove(ctx context.Context, id uint) error
}

type Storage interface {
	Save(ctx context.Context ,key string, file io.Reader) ( error)
	Remove(ctx context.Context, path, fileName string) error
	Open(ctx context.Context, key string) (io.ReadCloser, error)
}

type Broker interface{
	Publish(ctx context.Context, event string, body []byte) error
}

type Service struct {
	imageRepo ImageRepository
	storage   Storage
	broker Broker
}

func New(imageRepo ImageRepository, storage Storage, broker Broker) Service {
	return Service{
		imageRepo: imageRepo,
		storage:   storage,
		broker: broker,
	}


}

func (s Service) Upload(ctx context.Context, req param.UploadImageRequest) (param.UploadImageResponse, error) {
	fileHeader := req.Image
	file, _ := fileHeader.Open()
	id := uuid.New()
	ext := filepath.Ext(fileHeader.Filename)
	fileName := id.String() + ext
	err := s.storage.Save(ctx, "originals/" + fileName, file)
	if err != nil {
		return param.UploadImageResponse{}, fmt.Errorf("can't save image")
	}


	imageEntity := entity.Image{
		OriginalName: fileHeader.Filename,
		OriginalKey: fileName,
		ThumbnailKey: "",
		ContentType: ext,
		Size: uint(fileHeader.Size),
		Status: constant.PROCESSING,

	}

	imageRepo, err := s.imageRepo.Save(ctx, imageEntity)
	if err != nil {
		sErr := s.storage.Remove(ctx, "originals/" ,fileName)
		if sErr != nil {
			return param.UploadImageResponse{},fmt.Errorf("error in remove file: %w", sErr)
		}
		return param.UploadImageResponse{}, fmt.Errorf("unexpected error:%w", err)
	}

	evt := event.ImageUploaded{ImageID: imageRepo.ID}
	body, err := json.Marshal(evt)
	if err != nil {
		return param.UploadImageResponse{}, err
	}

	rErr := s.broker.Publish(ctx, "image_processing", body)
	if rErr != nil {
		return param.UploadImageResponse{}, rErr
	}



	return param.UploadImageResponse{
		ID: imageRepo.ID,
		OriginalName: fileHeader.Filename,
		OriginalKey: fileName,
		ThumbnailKey: imageRepo.ThumbnailKey,
		ContentType: ext,
		Size: uint(fileHeader.Size),
		Status: imageRepo.Status,

	}, nil
}

func (s Service) Get(ctx context.Context, req param.GetImageRequest) (param.GetImageResponse, error) {
	image, err := s.imageRepo.GetByID(ctx, req.ID)
	if err != nil {
		return param.GetImageResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return param.GetImageResponse{
		ID: image.ID,
		OriginalName: image.OriginalName,
		OriginalKey: image.OriginalKey,
		ThumbnailKey: image.ThumbnailKey,
		ContentType: image.ContentType,
		Size: image.Size,
		Status: image.Status,
	}, nil
}

func (s Service) DownloadOriginal(ctx context.Context, req param.DownloadImageRequest) (param.DownloadImageResponse, error) {
	image, err := s.imageRepo.GetByID(ctx, req.ID)
	if err != nil {
		return param.DownloadImageResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	file, err := s.storage.Open(ctx, "originals/" + image.OriginalKey)

	return param.DownloadImageResponse{
		ID: image.ID,
		OriginalName: image.OriginalName,
		OriginalKey: image.OriginalKey,
		ThumbnailKey: image.ThumbnailKey,
		ContentType: image.ContentType,
		Size: image.Size,
		Status: image.Status,
		File: file,
	}, nil

}

func (s Service) DownloadThumbnail(ctx context.Context, req param.DownloadImageRequest) (param.DownloadImageResponse, error) {
	image, err := s.imageRepo.GetByID(ctx, req.ID)
	if err != nil {
		return param.DownloadImageResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	file, err := s.storage.Open(ctx, "thumbnails/" + image.ThumbnailKey)

	return param.DownloadImageResponse{
		ID: image.ID,
		OriginalName: image.OriginalName,
		OriginalKey: image.OriginalKey,
		ThumbnailKey: image.ThumbnailKey,
		ContentType: image.ContentType,
		Size: image.Size,
		Status: image.Status,
		File: file,
	}, nil

}

func (s Service) RemoveImage(ctx context.Context, id uint) error {
	image, err := s.imageRepo.GetByID(ctx, id)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("unexpected error: %w", err)
	}

	oErr := s.storage.Remove(ctx, "originals/", image.OriginalKey )
	if oErr != nil {
		fmt.Println(oErr)
		return fmt.Errorf("unexpected error: %w", err)
	}

	tErr := s.storage.Remove(ctx, "thumbnails/", image.ThumbnailKey)
	if tErr != nil {
		fmt.Println(tErr)
		return fmt.Errorf("unexpected error: %w", tErr)
	}

	pErr := s.imageRepo.Remove(ctx, id)
	if pErr != nil {
		fmt.Println(pErr)
		return fmt.Errorf("unexpected error: %w", pErr)
	}

	return nil

}


func (s Service) ProcessImage(ctx context.Context, imageID uint) error {
	image, err := s.imageRepo.GetByID(ctx, imageID)

	if err != nil {
		return err
	}

	file, err := s.storage.Open(ctx, "originals/" + image.OriginalKey)


	img, err := imaging.Decode(file)
	if err != nil {
		return err
	}

	thumbnail := imaging.Thumbnail(img, 300, 300, imaging.Lanczos)

	buf := new(bytes.Buffer)
	err = imaging.Encode(buf, thumbnail, imaging.PNG)

	if err != nil {
		return nil
	}

	thumbnailKey := fmt.Sprintf("%d.png", image.ID)
	err = s.storage.Save(ctx, "thumbnails/" + thumbnailKey, buf)
	if err != nil {
		return err
	}

	image.Status = constant.COMPLETED
	image.ThumbnailKey = thumbnailKey

	err = s.imageRepo.Update(ctx, image)




	return nil
}