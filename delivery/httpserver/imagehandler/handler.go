package imagehandler

import "image_processing/service/imageservice"

type Handler struct {
	imageSvc imageservice.Service
}

func New(imageSvc imageservice.Service) Handler {
	return Handler{
		imageSvc: imageSvc,
	}
}