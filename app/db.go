package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (s *server) initDatabase(ctx context.Context) {
	dsn := "host=localhost user=postgres password=qwerty123 dbname=klikcair port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to open database connection:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("failed to get underlying sql.DB:", err)
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		// Tangani error timeout secara spesifik.
		if errors.Is(err, context.DeadlineExceeded) {

			// Dapatkan waktu deadline dari context.
			deadline, ok := ctx.Deadline()

			// Jika 'ok' bernilai true, context memiliki deadline.
			if ok {
				// Hitung durasi yang tersisa.
				timeoutRemaining := time.Until(deadline)
				fmt.Printf("Database connection has a deadline. Time remaining: %s\n", timeoutRemaining)
				log.Fatalln("database connection timed out after %v: %w", deadline, err)

				// Anda bisa menggunakan nilai ini dalam logika Anda.
				// Misalnya, untuk logger:
				// logger.Info("Connecting to DB with timeout", zap.Duration("timeout_remaining", timeoutRemaining))
			} else {
				fmt.Println("No deadline set for this context.")
			}

		}
		log.Fatalln("failed to ping database:", err)
	}

	s.db = db
	s.sqlDB = sqlDB
}
