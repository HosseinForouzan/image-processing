package imagehandler

import (
	"image_processing/service/imageservice"
	"image_processing/validator/imagevalidator"
)

type Handler struct {
	imageSvc imageservice.Service
	imageValidator imagevalidator.Validator
}

func New(imageSvc imageservice.Service) Handler {
	return Handler{
		imageSvc: imageSvc,
	}
}