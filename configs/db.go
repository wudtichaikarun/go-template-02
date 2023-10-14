package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host    string
	Port    string
	User    string
	Pass    string
	Name    string
	SSLMode string
}

func DBConnection() string {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.pass")
	name := viper.GetString("database.name")
	sslMode := viper.GetString("database.sslMode")
	engine := viper.GetString("database.engine")

	if engine == "postgres" {
		return fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			host, user, password, name, port, sslMode,
		)
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		user, password, host, port, name,
	)

}
