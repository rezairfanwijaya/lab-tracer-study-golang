package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type responseAPI struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func ResponseAPI(status string, code int, data interface{}) responseAPI {
	var meta meta
	meta.Code = code
	meta.Status = status

	return responseAPI{
		Meta: meta,
		Data: data,
	}
}

func GetENV() (map[string]string, error) {
	// read env file
	env, err := godotenv.Read(".env")
	if err != nil {
		return env, err
	}

	return env, nil
}

func ErrorBinding(err error) []string {
	var myError []string
	for _, e := range err.(validator.ValidationErrors) {
		myError = append(myError, e.Error())
	}

	return myError
}
