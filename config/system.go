package config

import (
	"log"

	"github.com/spf13/viper"
)

type EmailConfig struct {
	Email    string
	Password string
	SmtpHost string
	SmtpPort string
}



func (c *EmailConfig) LoadConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	c.Email = viper.GetString("EMAIL")
	c.Password = viper.GetString("PASSWORD")
	c.SmtpHost = viper.GetString("HOST")
	c.SmtpPort = viper.GetString("PORT")
}
