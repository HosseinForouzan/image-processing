package consumer

import (
	"context"
	"encoding/json"
	"image_processing/event"
)

func (c ImageCunsumer) Handle(ctx context.Context, body []byte) error {
	var evt event.ImageUploaded

	if err := json.Unmarshal(body, &evt); err !=nil {
		return err
	}

	return c.service.ProcessImage(ctx, evt.ImageID)
}

