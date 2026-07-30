package main

import (
	"fmt"
	"image_processing/broker"
	"image_processing/delivery/httpserver"
	"image_processing/repository/psql"
	"image_processing/repository/psql/psqlimage"
	"image_processing/repository/storage"
	"image_processing/service/imageservice"
	"log"
)

func main() {

	rabbitURL := "amqp://guest:guest@localhost:5672/"
	brokerRepo, err := broker.New(rabbitURL)
	if err != nil {
		log.Fatal(err)
	}
	defer brokerRepo.Close()

	fmt.Println("rabbit connected!")

	// _, err = mq.DeclareQueue("image_processing")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	root := "storage/originals"
	psql := psql.NewPgxPool()
	psqlImage := psqlimage.New(psql)

	storageRepo := storage.New(root)

	imageSvc := imageservice.New(psqlImage, storageRepo, brokerRepo)

	handler := httpserver.New(imageSvc)

	handler.Serve()
}