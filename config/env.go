package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost  string
	DBUser   string
	DBPasswd string
	DBPort   string
	DBName string
	Port string
	JwtSecret string
	JwtExpirationTime int
}

var ENV = NewConfig()

func NewConfig() *Config {
	godotenv.Load()
	return &Config{
		DBUser : getEnv("DB_USER","root"),
		DBPasswd: getEnv("DB_PASSWD","password"),
		DBPort: getEnv("DB_ADDR","3307"),
		DBName: getEnv("DB_NAME","jwt-go"),
		Port: getEnv("DB_HOST","5001"),
		DBHost: getEnv("DB_HOST","localhost"),
		JwtSecret: getEnv("JWT_SECRET","secret"),
		JwtExpirationTime:  getEnvInt("JWT_EXPIRATION_TIME",5),
	}
}

func getEnv(key, fallback string) string {
	if value,ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value,ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil{
			return fallback
		}
		return i
	}
	return fallback
}
