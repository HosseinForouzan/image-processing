package httpserver

import (
	"context"
	"fmt"
	"image_processing/delivery/httpserver/imagehandler"
	"image_processing/service/imageservice"
	"time"

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

func (s Server) Serve(ctx context.Context) error {
	s.Router.Use(middleware.Recover())
	s.Router.GET("/health-check", s.healthCheck)
	s.handler.SetRoutes(s.Router)


	address := fmt.Sprintf(":%d", 8080)
	fmt.Printf("start echo server on %s\n", address)

	sc := echo.StartConfig{
		Address: address,
		GracefulTimeout: 10 * time.Second,
	}

	return sc.Start(ctx, s.Router)

	// if err := s.Router.Start(address); err != nil {
	// 	fmt.Println("router start error", err)
	// }
}