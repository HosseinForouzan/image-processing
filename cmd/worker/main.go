package main

import (
	"context"
	"image_processing/broker"
	"image_processing/consumer"
	"image_processing/repository/psql"
	"image_processing/repository/psql/psqlimage"
	"image_processing/repository/storage"
	"image_processing/service/imageservice"
	"log"
)

func main() {
	ctx := context.Background()

	rabbitURL := "amqp://guest:guest@localhost:5672/"

	mq, err := broker.New(rabbitURL)
	if err != nil {
		log.Fatal(err)
	}
	defer mq.Close()

	_, err = mq.DeclareQueue("image_processing")
	if err != nil {
		log.Fatal(err)
	}

	root := "storage/"
	psql := psql.NewPgxPool()
	psqlImage := psqlimage.New(psql)

	storageRepo := storage.New(root)

	imageSvc := imageservice.New(psqlImage, storageRepo, mq)

	imageConsumer := consumer.New(imageSvc)


	log.Println("Worker Started...")



	err = mq.Consume(
		ctx,
		"image_processing",
		func(body []byte) error {
			return imageConsumer.Handle(ctx, body)
		},
	)

	if err != nil {
		log.Fatal(err)
	}

}