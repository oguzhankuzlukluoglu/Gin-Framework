package utils

import "github.com/spf13/viper"

// Config struct
type Config struct {
	DBName     string
	DBUser     string
	DBPassword string
	Secret     string
	Port       string
}

var config *Config

// GetConfig func
func GetConfig() *Config {
	return config
}

// LoadConfigVars -- load confiuration
func LoadConfigVars() (*Config, error) {
	// Set environment file
	viper.SetConfigFile(".env")

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		return new(Config), err
	}

	config = &Config{
		DBName:     viper.GetString("DB_NAME"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		Secret:     viper.GetString("SECRET"),
		Port:       viper.GetString("PORT"),
	}

	return config, nil
}
