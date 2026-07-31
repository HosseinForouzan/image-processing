package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

)

func (l LocalStorage) Save(ctx context.Context , key string, file io.Reader) error {

	path := filepath.Join(l.root, key)

	dst, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("can't create image file by the path: %w", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return fmt.Errorf("can't copy image to the file: %w", err)
	}

	return nil
}

func (l LocalStorage) Remove(ctx context.Context, path, fileName string) error {
	return os.Remove(l.root + path + fileName)
}

func (l LocalStorage) Open(ctx context.Context, key string) (io.ReadCloser, error) {
	path := filepath.Join(l.root, key)

	return os.Open(path)
}

