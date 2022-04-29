package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type GetEnv func(key string ) string

func LoadEnvironmentVariables() GetEnv {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	return os.Getenv
}
