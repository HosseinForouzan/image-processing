package imagehandler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


// RemoveImage deletes an image by ID.
//
// @Summary Delete image
// @Description Deletes an image along with its original and thumbnail files
// @Tags Images
// @Accept json
// @Produce json
// @Param id path int true "Image ID"
// @Success 204 "Image deleted successfully"
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Router /images/{id} [delete]
func (h Handler) Remove(c echo.Context) error {
	i := c.Param("id")
	id, err := strconv.Atoi(i)
	idUint := uint(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.imageSvc.RemoveImage(c.Request().Context(), idUint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusNoContent, "")
}