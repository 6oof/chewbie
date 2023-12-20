package helpers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Env retrieves the value of an environmental variable.
// If the variable is not set, it returns the default value.
//
// Example Usage:
//
//	value := Env("KEY_NAME", "default_value")
func Env(key, defaultValue string) string {

	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// EnvOrPanic retrieves the value of an environmental variable.
// If the variable is not set, it panics with an error message.
//
// Example Usage:
//
//	value := EnvOrPanic("KEY_NAME")
func EnvOrPanic(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic(fmt.Sprintf("Environmental variable %s is not set", key))
}

// LoadEnv loads the .env file.
//
// Example Usage:
//
//	LoadEnv()
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
