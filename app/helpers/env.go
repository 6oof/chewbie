package helpers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// GetEnvVar retrieves the value of an environmental variable.
// If the variable is not set, it returns the default value.
func Env(key, defaultValue string) string {

	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetEnvVarOrPanic retrieves the value of an environmental variable.
// If the variable is not set, it panics with an error message.
func EnvOrPanic(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic(fmt.Sprintf("Environmental variable %s is not set", key))
}

// Load the .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
