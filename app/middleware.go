// app/middleware.go
package app // Or whatever your package name is for this file

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware" // Alias for Echo's built-in middleware

	// Import your custom middleware packages with clear aliases
	customMiddleware "wallet-ex1/app/middleware"
)

func (s *server) initMiddleware() {

	// Pre-routing middleware
	s.server.Pre(
		middleware.RequestID(),
		middleware.Decompress(), // Dekompresi payload sebelum routing
	)

	s.server.Use(
		// Middleware untuk menangani CORS
		middleware.CORS(),
		// Middleware Recover untuk menangkap panic
		middleware.RecoverWithConfig(middleware.RecoverConfig{
			Skipper: func(c echo.Context) bool {
				return false
			},
			StackSize:           4 << 10, // 4 KB
			DisableStackAll:     true,
			DisablePrintStack:   false,
			LogLevel:            0,
			DisableErrorHandler: false,
		}),

		// Use your custom logger middleware with its correct alias
		customMiddleware.LoggerMiddleware(customMiddleware.LoggerConfig{
			Logger: s.log, // Pass the initialized server logger
		}),

		// Use your custom DB middleware with its correct alias
		customMiddleware.DBMiddleware(s.db),
	)
}
