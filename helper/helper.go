package helper

import (
	"github.com/joho/godotenv"
)

func GetENV() (map[string]string, error) {
	// read env file
	env, err := godotenv.Read(".env")
	if err != nil {
		return env, err
	}

	return env, nil
}
