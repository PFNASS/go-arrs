package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
	KeyLookup: "query:api-key",
	Validator: func(key string, c echo.Context) (bool, error) {
				return key == "valid-key", nil
			},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.Start(":8080")
}