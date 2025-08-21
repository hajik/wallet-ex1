package middleware

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Middleware database.
func DBMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Tempatkan instance DB di konteks Echo.
			c.Set("db", db)
			return next(c)
		}
	}
}
