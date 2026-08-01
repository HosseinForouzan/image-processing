// @title           Image Processing API
// @version         1.0
// @description     Image Processing Service with RabbitMQ
// @BasePath        /

// @contact.name    Hossein Forouzan
// @contact.email   hosseinf97@gmail.com

package main


import (
	"context"
	"fmt"
	"image_processing/broker"
	"image_processing/delivery/httpserver"
	"image_processing/repository/psql"
	"image_processing/repository/psql/psqlimage"
	"image_processing/repository/storage"
	"image_processing/service/imageservice"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	root := "storage/"
	psql := psql.NewPgxPool()
	psqlImage := psqlimage.New(psql)

	storageRepo := storage.New(root)

	imageSvc := imageservice.New(psqlImage, storageRepo, brokerRepo)

	server := httpserver.New(imageSvc)
	server.Serve()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	fmt.Println("Shutting down the server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	if err := server.Router.Shutdown(shutdownCtx); err != nil {
		log.Println("server shutdown error:", err)
	}

	fmt.Println("server exited succesfully!")

}