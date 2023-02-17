package main
import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo-contrib/prometheus"
)
func main() {
    e := echo.New()
    // Enable metrics middleware
    p := prometheus.NewPrometheus("echo", nil)
    p.Use(e)

    e.Logger.Debug(e.Start(":1323"))
}