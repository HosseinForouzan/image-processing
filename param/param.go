package param
import "mime/multipart"

type UploadImageRequest struct {
	Image *multipart.FileHeader `json:"image"`
}

type UploadImageResponse struct {
	ID uint `json:"id"`
	OriginalName string `json:"original_name"`
	OriginalKey string `json:"original_key"`
	ThumbnailKey string `json:"thumbnail_key"`
	ContentType string `json:"content_type"`
	Size uint `json:"size"`
	Status string `json:"status"`
}

type GetImageRequest struct {
	ID uint `josn:"id"`
}

type GetImageResponse struct {
	ID uint `json:"id"`
	OriginalName string `json:"original_name"`
	OriginalKey string `json:"original_key"`
	ThumbnailKey string `json:"thumbnail_key"`
	ContentType string `json:"content_type"`
	Size uint `json:"size"`
	Status string `json:"status"`
}