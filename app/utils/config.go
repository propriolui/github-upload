package utils

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Configurations : contiene tutte le configurazioni utili al server e jwt
type Configurations struct {
	ServerAddress           string
	DBurl                   string
	AccessTokenSecret       string
	RefreshTokenSecret      string
	JwtExpiration           int // in minutes
	MailVerifCodeExpiration int // in hours
	PassResetCodeExpiration int // in minutes
	MailVerifTemplateID     string
	PassResetTemplateID     string
}

//NewConfigurations : ritorna un nuovo oggetto configurazione
func NewConfigurations(logger *zap.SugaredLogger) *Configurations {
	viper.AutomaticEnv()

	configs := &Configurations{
		ServerAddress:      viper.GetString("SERVER_ADDRESS"),
		DBurl:              viper.GetString("DB_URL"),
		AccessTokenSecret:  viper.GetString("ACCESS_SECRET"),
		RefreshTokenSecret: viper.GetString("REFRESH_SECRET"),
	}
	return configs
}
