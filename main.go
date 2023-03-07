package main
import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo-contrib/prometheus"
    "net/http"
)
func main() {
    e := echo.New()
    // Enable metrics middleware
    e.GET("/users/:name", func(c echo.Context) error {
        name := c.Param("name")
        return c.String(http.StatusOK, name)
    })

    e.GET("/surote-cal", func(c echo.Context) error {
        return c.String(http.StatusOK, "surote-cal is OK~")
    })
    p := prometheus.NewPrometheus("echo", nil)
    p.Use(e)

    e.Logger.Debug(e.Start(":1323"))
}