package env

import (
	"os"
)

//GetVariable returns .env variable value
func GetVariable(key string) string {
	return os.Getenv(key)
}
