package controller

import(
	"net/http"
	"github.com/labstack/echo/v4"
	"app/datastore"
	"context"
)

func GetBooks() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		var err error
		client := datastore.GetConnection()
		type Data struct {
			Name string
		}


		posts := make([]*Data, 0)
		_, err = client.GetAll(ctx, client.NewQuery("Data").Order("Name"), &posts)
		if err != nil {
		    panic(err)
		}


		return c.JSON(http.StatusOK, posts)
	}
}