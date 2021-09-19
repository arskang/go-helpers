package goHelpers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetENV(name string) (*string, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("the .env file could not be loaded, %v", err.Error())
	}

	env := os.Getenv(name)
	if env == "" {
		return nil, fmt.Errorf("environment variable not found")
	}

	return &env, nil
}
