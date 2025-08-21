package app

import (
	"context"
	"database/sql"
	"errors"
	stdLog "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// deadlineTask set maximum timeout for each setup task
const deadlineTask = 30

type server struct {
	server *echo.Echo
	db     *gorm.DB
	sqlDB  *sql.DB
	log    *zap.Logger
}

// Run setup new http server that use echo under the hood then run it.
func Run() {

	// Create a new context for the shutdown with a specific deadline
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	serv := &server{server: echo.New()}
	serv.server.HideBanner = true
	serv.server.Debug = false

	// Trigger some initialization
	serv.initAppConfig()
	serv.initLogger()
	defer serv.log.Sync()

	serv.initDatabase(shutdownCtx)
	serv.initValidator()
	serv.initMiddleware()
	serv.initRoutes()

	// Run the HTTP server on a separate goroutine
	go func() {
		// Start server with port from config
		host := viper.GetString("app.host")
		port := viper.GetString("app.port")
		if err := serv.server.Start(host + ":" + port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			stdLog.Fatalf("Failed to start the server: %v", err)
		}
	}()

	// Create channel that listens for termination signals
	quit := make(chan os.Signal, 1)
	// Sinyal SIGKILL tidak dapat dicegat, jadi jangan sertakan
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block the program until terminate signal received
	<-quit

	// --- Graceful Shutdown ---
	stdLog.Println("Gracefully shutting down...")

	// Shutdown http server
	if err := serv.server.Shutdown(shutdownCtx); err != nil {
		stdLog.Printf("Server forced to shutdown: %v", err)
	}

	// Close database connection
	stdLog.Println("Closing database connection...")
	if serv.db != nil {
		sqlDb, err := serv.db.DB()
		if err != nil {
			stdLog.Printf("Failed to get database driver: %v", err)
		}
		if err = sqlDb.Close(); err != nil {
			stdLog.Printf("Failed to close database: %v", err)
		}
	}

	stdLog.Println("Server Exiting...")
	os.Exit(0) // Exit with code 0 to indicate a clean exit
}
