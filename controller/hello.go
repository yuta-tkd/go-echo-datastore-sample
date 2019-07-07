package controller

import(
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetHello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
	}
}