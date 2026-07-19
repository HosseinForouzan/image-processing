package imagehandler

import (
	"net/http"
	"github.com/labstack/echo/v5"
)

func (h Handler) Upload(c *echo.Context) error {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.imageSvc.Storage.Save(fileHeader)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}



	

	// return c.JSON(http.StatusCreated, map[string]any{
	// 	"filename": fileHeader.Filename,
	// 	"size": fileHeader.Size,
	// 	"content-type": fileHeader.Header.Get("Content-Type"),
	// })

	return c.JSON(http.StatusCreated, resp)


}