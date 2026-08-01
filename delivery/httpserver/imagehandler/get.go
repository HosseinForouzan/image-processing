package imagehandler

import (
	"image_processing/param"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


// GetImage returns image information by ID.
//
// @Summary Get image information
// @Description Returns metadata of an image by its ID
// @Tags Images
// @Accept json
// @Produce json
// @Param id path int true "Image ID"
// @Success 200 {object} param.GetImageResponse
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Router /images/{id} [get]
func (h Handler) Get(c echo.Context) error {
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