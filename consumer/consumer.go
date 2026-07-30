package consumer

import "image_processing/service/imageservice"

type ImageCunsumer struct {
	service imageservice.Service
}

func New(service imageservice.Service) ImageCunsumer {
	return ImageCunsumer{
		service: service,
	}
}

