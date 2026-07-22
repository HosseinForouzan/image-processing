package param
import "mime/multipart"

type SaveImageRequest struct {
	FileHeader *multipart.FileHeader
}

type SaveImageResponse struct {
	ID uint
	OriginalName string
	OriginalKey string
	ThumbnailKey string
	ContentType string
	Size uint
	Status string
}