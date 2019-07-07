package controller

import(
	"net/http"
	"github.com/labstack/echo/v4"
	"app/datastore"
	"context"
	"time"
	"strconv"
)


func GetBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		bookIdStr := c.Param("book_id")
		bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
		if err != nil {
			panic(err)
		}

		ctx := context.Background()
		client := datastore.GetConnection()


		key := client.IDKey("Book", bookId,nil)
		entity := &datastore.Book{}
		err = client.Get(ctx, key, entity)
		if err != nil {
			panic(err)
		}

		res := entity

		return c.JSON(http.StatusOK,res)
	}
}

func PostBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		client := datastore.GetConnection()

		var err error

		title := c.QueryParam("title")
		if title == "" {
			panic(err)
		}


		key := client.IncompleteKey("Book", nil)
		entity := &datastore.Book{
			Title: title,
			TimeStamp: time.Now().Unix(),
		}
		key, err = client.Put(ctx, key, entity)
		if err != nil {
			panic(err)
		}

		res := map[string]int64{
	        "key": key.ID(),
	    }

		return c.JSON(http.StatusOK,res)
	}
}


func PutBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		bookIdStr := c.Param("book_id")
		bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
		if err != nil {
			panic(err)
		}

		ctx := context.Background()
		client := datastore.GetConnection()


		key := client.IDKey("Book", bookId,nil)
		entity := &datastore.Book{}
		err = client.Get(ctx, key, entity)
		if err != nil {
			panic(err)
		}


		title := c.QueryParam("title")
		if title == "" {
			panic(err)
		}
		entity.Title = title


		key, err = client.Put(ctx, key, entity)
		if err != nil {
			panic(err)
		}

		res := entity

		return c.JSON(http.StatusOK,res)
	}
}