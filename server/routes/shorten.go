package routes

import (
	"net/http"
	"net/url"
	"time"

	"github.com/Dank-del/link-shortner-api/types"

	"github.com/Dank-del/link-shortner-api/database"
	"github.com/labstack/echo/v4"
)

func Shorten(c echo.Context) error {
	uri := c.QueryParam("url")
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.Error{ErrData: err, OccouredAt: time.Now()})
	}
	data := database.CreateNewShortenedLink(uri)
	return c.JSON(http.StatusOK, data)
}
