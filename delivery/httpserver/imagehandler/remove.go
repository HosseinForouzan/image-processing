package imagehandler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func (h Handler) Remove(c *echo.Context) error {
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