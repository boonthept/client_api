//simple_api.go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		t := time.Now()
		fmt.Println(t.Format("20060102150405"))
		fmt.Println("... call to client ")
		return c.String(http.StatusOK, "Hello, world")
	})
	e.Logger.Fatal(e.Start(":8888"))
}
