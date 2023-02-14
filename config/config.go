package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Print("Error loading .env file")
		return "nil"
	}

    return os.Getenv(key)
}	


// func LoadConfigs() map[string]string {

// 	configs := map[string]string{
// 		"FRONT_END_PORT": getEnvValue("FRONT_END_PORT"),
// 		"BACK_END_PORT": getEnvValue("BACK_END_PORT"),

// 		"DB_HOST": getEnvValue("DB_HOST"),
// 		"DB_NAME": getEnvValue("DB_NAME"),
// 		"DB_USER": getEnvValue("DB_USER"),
// 		"DB_PASSWORD": getEnvValue("DB_PASSWORD"),
// 		"DB_PORT": getEnvValue("DB_PORT"),
// 	}
	
// 	return configs
// }