package imagehandler

import (

	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func (h Handler) Retry(c *echo.Context) error {
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