package imagehandler

import "github.com/labstack/echo/v5"

func (h Handler) SetRoutes(e *echo.Echo) {
	imageGroup := e.Group("/images")

	imageGroup.POST("/upload", h.Upload)
	imageGroup.GET("/:id", h.Get)
	imageGroup.GET("/:id/original", h.DownloadOriginal)

}

