package main

import (
	"image_processing/delivery/httpserver"
	"image_processing/repository/psql"
	"image_processing/repository/psql/psqlimage"
	"image_processing/repository/storage"
	"image_processing/service/imageservice"
)

func main() {
	root := "storage/originals"
	psql := psql.NewPgxPool()
	psqlImage := psqlimage.New(psql)

	storageRepo := storage.New(root)

	imageSvc := imageservice.New(psqlImage, storageRepo)

	handler := httpserver.New(imageSvc)

	handler.Serve()
}