package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manimovassagh/htmx-g-tmpl/models"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		// Return the First user as JSON data
		return c.JSON(http.StatusOK, models.First)
	})

	e.Logger.Fatal(e.Start(":8080"))

}
