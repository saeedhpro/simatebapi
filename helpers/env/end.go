package env

import "os"

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
