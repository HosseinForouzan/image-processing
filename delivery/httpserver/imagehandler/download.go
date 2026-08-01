package imagehandler

import (
	"fmt"
	"image_processing/param"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// DownloadOriginal downloads the original image.
//
// @Summary Download original image
// @Description Download original image by ID
// @Tags Images
// @Produce application/octet-stream
// @Param id path string true "Image ID"
// @Success 200 {file} param.DownloadImageRequest
// @Failure 404 {object} error
// @Router /images/{id}/original [get]
func (h Handler) DownloadOriginal(c echo.Context) error {
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


// DownloadThumbnail downloads the thmbnail image.
//
// @Summary Download thumbnail image
// @Description Download thumbnail image by ID
// @Tags Images
// @Produce application/octet-stream
// @Param id path string true "Image ID"
// @Success 200 {file} param.DownloadImageRequest
// @Failure 404 {object} error
// @Router /images/{id}/thumbnail [get]
func (h Handler) DownloadThumbnail(c echo.Context) error {
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