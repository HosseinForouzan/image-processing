package imagehandler

import (
	"image_processing/param"
	"net/http"

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

