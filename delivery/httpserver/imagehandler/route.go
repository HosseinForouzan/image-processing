package imagehandler

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "image_processing/docs"

)

func (h Handler) SetRoutes(e *echo.Echo) {

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	imageGroup := e.Group("/images")

	imageGroup.POST("/upload", h.Upload)
	imageGroup.GET("/:id", h.Get)
	imageGroup.GET("/:id/original", h.DownloadOriginal)
	imageGroup.GET("/:id/thumbnail", h.DownloadThumbnail)
	imageGroup.DELETE("/:id", h.Remove)
	imageGroup.POST("/:id/retry", h.Retry)
	



}

