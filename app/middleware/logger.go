// app/middleware/logger.go
package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// LoggerConfig is as you defined
type LoggerConfig struct {
	Logger *zap.Logger
}

// LoggerMiddleware is as you defined
func LoggerMiddleware(config LoggerConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			c.Set("logger", config.Logger) // This correctly sets the logger in context

			err := next(c)

			req := c.Request()
			res := c.Response()
			stop := time.Now()

			// Prepare common fields to avoid repetition
			fields := []zap.Field{
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.String("remote_ip", c.RealIP()),
				zap.Int("status", res.Status),
				zap.Duration("latency", stop.Sub(start)),
			}

			if err != nil {
				// Add error field for failed requests
				fields = append(fields, zap.Error(err))
				config.Logger.Error("Request Failed", fields...)
				c.Error(err)
			} else {
				// Add response size for successful requests
				fields = append(fields, zap.Int64("response_size", res.Size))
				config.Logger.Info("Request Handled", fields...)
			}

			return err
		}
	}
}
