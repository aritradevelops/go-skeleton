package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName string
	Http        HttpConfig
	Db          DatabaseConfig
}

type HttpConfig struct {
	Host string
	Port int
}

type DatabaseConfig struct {
	Url string
}

func Load(readEnv bool) (Config, error) {
	var config Config
	// load environment variables from .env file
	if readEnv {
		err := godotenv.Load()
		if err != nil {
			return config, fmt.Errorf("failed to load environment variables: %v", err)
		}
	}

	httpConfig := HttpConfig{
		Host: getEnv("HTTP_HOST"),
		Port: getEnvInt("HTTP_PORT"),
	}

	dbConfig := DatabaseConfig{
		Url: getEnv("DATABASE_URL"),
	}

	config.Http = httpConfig
	config.Db = dbConfig
	config.ServiceName = getEnv("SERVICE_NAME")

	return config, nil
}

func getEnv(name string) string {
	return os.Getenv(name)
}

func getEnvInt(name string) int {
	val, _ := strconv.Atoi(getEnv(name))
	return val
}
