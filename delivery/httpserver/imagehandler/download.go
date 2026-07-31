package imagehandler

import (
	"fmt"
	"image_processing/param"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func (h Handler) DownloadOriginal(c *echo.Context) error {
	i := c.Param("id")
	id, err := strconv.Atoi(i)
	idUint := uint(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := param.DownloadImageRequest{ID: idUint}
	
	resp, err := h.imageSvc.DownloadOriginal(c.Request().Context(), req)

	c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf(`attachment; filename="%s"`, resp.OriginalName))

	return c.Stream(http.StatusOK, resp.ContentType, resp.File)
}

func (h Handler) DownloadThumbnail(c *echo.Context) error {
	i := c.Param("id")
	id, err := strconv.Atoi(i)
	idUint := uint(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := param.DownloadImageRequest{ID: idUint}
	
	resp, err := h.imageSvc.DownloadThumbnail(c.Request().Context(), req)

	c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf(`attachment; filename="%s"`, resp.ThumbnailKey))

	return c.Stream(http.StatusOK, resp.ContentType, resp.File)
}