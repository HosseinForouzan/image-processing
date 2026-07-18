package main

import (
	"fmt"
	"image_processing/delivery/httpserver"
	"image_processing/repository/psql"
	"image_processing/service/imageservice"
)

func main() {
	p := psql.NewPgxPool()
	fmt.Println(p)

	imageSvc := imageservice.New()

	handler := httpserver.New(imageSvc)

	handler.Serve()
}