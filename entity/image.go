package entity

import "image_processing/constant"

type Image struct {
	ID           uint
	OriginalName string
	OriginalKey  string
	ThumbnailKey string
	ContentType  string
	Size         uint
	Status       string
}

func (i Image) CanRetry() bool {
	return i.Status == constant.FAILED
}