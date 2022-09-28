package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Body struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func main() {
	app := echo.New()
	app.GET("/", func(ctx echo.Context) error {
		body := new(Body)

		if err := ctx.Bind(body); err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{
				"msg": "Bad request",
			})
		}

		fmt.Println(body)

		return ctx.JSON(http.StatusOK, body)
	})

	app.Logger.Fatal(app.Start(":8080"))
}
