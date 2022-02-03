package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

// Courtesy: https://www.wolfe.id.au/2020/03/10/how-do-i-structure-my-go-project/
// Code: https://echo.labstack.com/guide/#hello-world
