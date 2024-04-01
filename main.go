package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {

	app := echo.New()
	//e.GET("/user", helloHandler)

	app.Start("8080")
	fmt.Println("Check")
}
