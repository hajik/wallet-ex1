package app

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// initLogger menginisialisasi Zap logger dengan konfigurasi yang bisa diubah.
func (s *server) initLogger() {
	// Konfigurasi logger untuk produksi
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder // Menggunakan format waktu standar
	jsonEncoder := zapcore.NewJSONEncoder(config)

	// Menggunakan stdout sebagai output
	core := zapcore.NewCore(jsonEncoder, os.Stdout, zapcore.InfoLevel)

	s.log = zap.New(core, zap.AddCaller())

}
