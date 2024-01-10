package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func createDummyEnvFile(dir string) error {
	// Create a temporary .env file with a dummy variable
	envFilePath := filepath.Join(dir, ".env")
	return os.WriteFile(envFilePath, []byte("DUMMY_VAR=test_value"), 0644)
}

func TestLoadEnv(t *testing.T) {
	// Create a dummy .env file in the current working directory
	err := createDummyEnvFile(".")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(".env") // Clean up after the test

	// Test the LoadEnv function
	LoadEnv()

	// Check if the dummy variable is loaded
	value := os.Getenv("DUMMY_VAR")
	if value != "test_value" {
		t.Errorf("LoadEnv did not load the expected value. Got: %s, Expected: test_value", value)
	}
}

func TestLoadEnvNoFile(t *testing.T) {
	// Test the LoadEnv function without creating a dummy .env file
	// This test should ensure that LoadEnv fails gracefully when .env is not present

	// Ensure LoadEnv panics with the expected error message
	defer func() {
		if r := recover(); r != nil {
			errMsg := fmt.Sprintf("%s", r)
			if errMsg != "open .env: no such file or directory" {
				t.Errorf("LoadEnv did not panic with the expected error message. Got: %s", errMsg)
			}
		} else {
			t.Error("LoadEnv did not panic as expected for missing .env file")
		}
	}()

	LoadEnv()
}

// env_test.go (continued)
func TestEnv(t *testing.T) {
	// Set a dummy environmental variable in the existing .env file
	os.Setenv("DUMMY_VAR", "test_value")

	// Test the Env function
	value := Env("DUMMY_VAR", "default_value")
	if value != "test_value" {
		t.Errorf("Env did not return the expected value. Got: %s, Expected: test_value", value)
	}

	// Test with a non-existing variable
	value = Env("NON_EXISTING_VAR", "default_value")
	if value != "default_value" {
		t.Errorf("Env did not return the expected default value for non-existing variable. Got: %s, Expected: default_value", value)
	}
}

func TestEnvOrPanic(t *testing.T) {
	// Set a dummy environmental variable in the existing .env file
	os.Setenv("DUMMY_VAR", "test_value")

	// Test the EnvOrPanic function
	value := EnvOrPanic("DUMMY_VAR")
	if value != "test_value" {
		t.Errorf("EnvOrPanic did not return the expected value. Got: %s, Expected: test_value", value)
	}

	// Test with a non-existing variable
	defer func() {
		// Recover from panic and check if the panic message contains the variable name
		if r := recover(); r != nil {
			errMsg := fmt.Sprintf("%s", r)
			if errMsg != "Environmental variable NON_EXISTING_VAR is not set" {
				t.Errorf("EnvOrPanic did not panic with the expected error message. Got: %s", errMsg)
			}
		} else {
			t.Error("EnvOrPanic did not panic as expected for non-existing variable")
		}
	}()

	EnvOrPanic("NON_EXISTING_VAR")
}
