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