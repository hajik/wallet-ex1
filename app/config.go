package app

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	_ "gopkg.in/ini.v1" // Penting: Ini mendaftarkan parser INI
)

func (s *server) initAppConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf/")

	// Attempt to read the config file.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error reading config file: %s", err)
	}

	fmt.Printf("DB connected with db_name: %s\n", viper.GetString("app.name"))
}
