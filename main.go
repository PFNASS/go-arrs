package main

import (
	"os"
	"net/http"
	"log/slog"
	"context"
	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/devopsarr/radarr-go/radarr"
	cfg "github.com/PFNASS/go-arrs/pkg/config"
	"github.com/PFNASS/go-arrs/pkg/plex"
)

func main() {
	e := echo.New()

	config, err := cfg.loadConfig()
	if err != nil {
		e.Logger.Fatal("Failed to load config:", err)
	}

	plex, err := plex.NewPlexClient(
		cfg.GetString("plex.host"),
		cfg.GetString("plex.apiKey"),
		cfg.GetString("plex.apiSecret"),
		cfg.GetDuration("plex.timeout"),
	)
	if err != nil {
		e.Logger.Fatal("Failed to create Plex client:", err)
	}
	

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}