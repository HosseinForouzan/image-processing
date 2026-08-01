

package imagehandler

import (
	"image_processing/param"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UploadImage godoc
//
// @Summary Upload image
// @Description Upload original image
// @Tags Images
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Image"
// @Success 201 {object} param.UploadImageResponse
// @Failure 400 {object} error
// @Router /images/upload [post]
func (h Handler) Upload(c echo.Context) error {
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

