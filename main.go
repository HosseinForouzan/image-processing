package main

import (
	"image_processing/delivery/httpserver"
	"image_processing/repository/psql"
	"image_processing/repository/storage"
	"image_processing/service/imageservice"
)

func main() {
	root := "storage/originals"
	psqlRepo := psql.NewPgxPool()
	storageRepo := storage.New(root)

	imageSvc := imageservice.New(psqlRepo, storageRepo)

	handler := httpserver.New(imageSvc)

	handler.Serve()
}