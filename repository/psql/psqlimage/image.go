package psqlimage

import (
	"context"
	"fmt"
	"image_processing/entity"

)

func (d *DB) Save(ctx context.Context, image entity.Image) (entity.Image, error) {
	var id uint
	query := `INSERT INTO images(original_name, original_key, thumbnail_key, content_type, size, status)
			VALUES($1, $2, $3, $4, $5, $6) RETURNING id`
	err := d.conn.Conn().QueryRow(ctx, query, image.OriginalName, image.OriginalKey,
		 image.ThumbnailKey, image.ContentType, image.Size, image.Status).Scan(&id)
	if err != nil {
		return entity.Image{}, fmt.Errorf("can't insert image:%w", err)
	}

	image.ID = id

	return image, nil
}

func (d *DB) GetByID(ctx context.Context, id uint) (entity.Image, error) {
	var image entity.Image

	query := `SELECT id,original_name,original_key,thumbnail_key, content_type,size,status FROM images WHERE id = $1`
	err := d.conn.Conn().QueryRow(ctx, query, id).Scan(&image.ID, &image.OriginalName,
		 &image.OriginalKey, &image.ThumbnailKey, &image.ContentType, &image.Size, &image.Status)
	if err != nil {
		return entity.Image{}, fmt.Errorf("can't get images: %w", err)
	}

	return image, nil
}

func (d *DB) Update(ctx context.Context, image entity.Image) error {
	query := `UPDATE images SET thumbnail_key = $1, status=$2, updated_at = NOW() WHERE id = $3`
	_, err := d.conn.Conn().Exec(ctx, query, image.ThumbnailKey, image.Status, image.ID)

	return err
}

func (d *DB) Remove(ctx context.Context, id uint) error {
	query := `DELETE FROM images WHERE id = $1`
	_, err := d.conn.Conn().Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("can't remove item:%w", err)
	}

	return nil
}