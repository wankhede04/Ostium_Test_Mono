package config

import "github.com/spf13/viper"

type DBConfig struct {
	Driver   string // DataBase driver
	Host     string // DataBase host
	Port     string
	Ssl      string // DataBase sslmode
	Name     string // DataBase name
	User     string // DataBase's user
	Password string // User's password
}

func ReadDBConfig() DBConfig {
	return DBConfig{
		Driver:   viper.GetString("db.driver"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Ssl:      viper.GetString("db.ssl_mode"),
		Name:     viper.GetString("db.db_name"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
	}
}
