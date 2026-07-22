package param
import "mime/multipart"

type UploadImageRequest struct {
	Image *multipart.FileHeader
}

type UploadImageResponse struct {
	ID uint
	OriginalName string
	OriginalKey string
	ThumbnailKey string
	ContentType string
	Size uint
	Status string
}