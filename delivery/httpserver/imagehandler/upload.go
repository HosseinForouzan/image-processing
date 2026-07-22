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
	req := param.SaveImageRequest{
		FileHeader: fileHeader,
	}

	resp, err := h.imageSvc.SaveImage(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}



	return c.JSON(http.StatusCreated, resp)


}