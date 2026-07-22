package imagevalidator

import (
	"fmt"
	"image_processing/param"
	"io"
	"mime/multipart"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)


func (v Validator) ValidateUploadImageRequest(req param.UploadImageRequest) (error) {
	err := validation.ValidateStruct(&req,
		validation.Field(
			&req.Image,
			validation.Required,
			validation.By(v.validateImageSize),
			validation.By(v.ValidateImageType),
		),
	
	)

	if err != nil {
		return err
	}

	return nil
}

func (v Validator) validateImageSize(value interface{}) error {
	file, ok := value.(*multipart.FileHeader)
	if !ok {
		return fmt.Errorf("Invalid file")
	}

	const maxSize = 10 * 1024 * 1024

	if file.Size > maxSize {
		return fmt.Errorf("image size must be less than 10 mb")
	}

	return nil
}

func(v Validator) ValidateImageType(value interface{}) error {
	fileHeader, ok := value.(*multipart.FileHeader)
	if !ok {
		return fmt.Errorf("Invalid file")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 512)

	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return err
	}

	contentType := http.DetectContentType(buffer)

	if contentType == "image/jpeg" || contentType == "image/png" || contentType == "image/webp" {
		return nil
	}

	return fmt.Errorf("unsupported image type.")
}