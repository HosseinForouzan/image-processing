package imagehandler

import (
	"image_processing/param"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func (h Handler) Upload(c *echo.Context) error {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	req := param.UploadImageRequest{
		Image: fileHeader,
	}

	if err := h.imageValidator.ValidateUploadImageRequest(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.imageSvc.Upload(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}



	return c.JSON(http.StatusCreated, resp)


}

func (h Handler) Get(c *echo.Context) error {
	i := c.Param("id")
	id, err := strconv.Atoi(i)
	idUint := uint(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := param.GetImageRequest{ID: idUint}

	resp, err := h.imageSvc.Get(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)

}