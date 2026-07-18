package httpserver

import (
	"fmt"
	"image_processing/delivery/httpserver/imagehandler"
	"image_processing/service/imageservice"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	handler imagehandler.Handler
	Router *echo.Echo

}

func New(imageSvc imageservice.Service) Server {
	return Server{
		handler: imagehandler.New(imageSvc),
		Router: echo.New(),
	}
}

func (s Server) Serve() {
	s.Router.Use(middleware.Recover())
	s.Router.GET("/health-check", s.healthCheck)


	address := fmt.Sprintf(":%d", 8080)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}