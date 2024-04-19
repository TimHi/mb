package server

import (
	"log/slog"

	"github.com/Pineapple217/mb/pkg/database"
	"github.com/Pineapple217/mb/pkg/handler"
	"github.com/Pineapple217/mb/pkg/middleware"
	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
)

func (s *Server) ApplyMiddleware(q *database.Queries) {
	s.e.Use(echoMw.RequestLoggerWithConfig(echoMw.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogMethod:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v echoMw.RequestLoggerValues) error {
			slog.Info("request",
				"method", v.Method,
				"status", v.Status,
				"latency", v.Latency,
				"path", v.URI,
			)
			return nil

		},
	}))

	s.e.Use(echoMw.GzipWithConfig(echoMw.GzipConfig{
		Level: 5,
	}))

	echo.NotFoundHandler = handler.NotFound
	s.e.Use(func(next echo.HandlerFunc) echo.HandlerFunc { return middleware.Stats(next, q) })
	s.e.Use(middleware.Auth)
}