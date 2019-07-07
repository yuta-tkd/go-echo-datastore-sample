package route

import(
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"app/controller"
)

func Init() *echo.Echo {

	e := echo.New()

	initMiddleware(e)
	initRouting(e)

	return e
}

func initMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func initRouting(e *echo.Echo) {
	e.GET("/", controller.GetHello())

	e.GET("/book/:book_id", controller.GetBook())
	e.PUT("/book/:book_id", controller.PutBook())
	e.POST("/book", controller.PostBook())

	e.GET("/books", controller.GetBooks())
}
