package env

import (
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func SetEnv(key string, val string) error {
	err := os.Setenv(key, val)
	if err != nil {
		return err
	}
	return nil
}
