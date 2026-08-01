package imagehandler

import (

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// RetryImageProcessing retries thumbnail generation for a failed image.
//
// @Summary Retry image processing
// @Description Republishes the image processing job to RabbitMQ
// @Tags Images
// @Accept json
// @Produce json
// @Param id path int true "Image ID"
// @Success 204 "Image queued for processing"
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Router /images/{id}/retry [post]
func (h Handler) Retry(c echo.Context) error {
	i := c.Param("id")
	id, err := strconv.Atoi(i)
	idUint := uint(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	
	err = h.imageSvc.Retry(c.Request().Context(), idUint)
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}


	return c.NoContent(http.StatusNoContent)
}