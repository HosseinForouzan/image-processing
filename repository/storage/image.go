package storage

import (
	"context"
	"fmt"
	"image_processing/entity"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func (l LocalStorage) Save(ctx context.Context,fileHeader *multipart.FileHeader) (entity.Image,error) {
	file, err := fileHeader.Open()
	if err != nil {
		return entity.Image{}, fmt.Errorf("error in open image: %w", err)
	}
	id := uuid.New()
	ext := filepath.Ext(fileHeader.Filename)
	fileName := id.String() + ext
	path := filepath.Join(l.root, fileName)

	dst, err := os.Create(path)
	if err != nil {
		return entity.Image{} ,fmt.Errorf("can't create image file by the path: %w", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return entity.Image{}, fmt.Errorf("can't copy image to the file: %w", err)
	}

	return entity.Image{
		OriginalName: fileHeader.Filename,
		OriginalKey: fileName,
		ContentType: ext,
		Size: uint(fileHeader.Size),

	}, nil
}

func (l LocalStorage) Remove(ctx context.Context, fileName string) error {
	return os.Remove(l.root + "/" + fileName)
}

func (l LocalStorage) Open(ctx context.Context, key string) (io.ReadCloser, error) {
	path := filepath.Join(l.root, key)

	return os.Open(path)
}

