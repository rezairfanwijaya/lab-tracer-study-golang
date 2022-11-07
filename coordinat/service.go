package coordinat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CoordinateService interface {
	GetCoordinate(city string) (coordinate, error)
}

type coordinateService struct{}

type coordinate struct {
	ResponseCode string `json:"responseCode"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
}

func NewCoordinateService() *coordinateService {
	return &coordinateService{}
}

func (c *coordinateService) GetCoordinate(city string) (coordinate, error) {
	URL := fmt.Sprintf("http://www.gps-coordinates.net/api/%s", city)
	resp, err := http.Get(URL)
	if err != nil {
		return coordinate{}, err
	}

	respbyte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return coordinate{}, err
	}

	var result coordinate
	if err := json.Unmarshal(respbyte, &result); err != nil {
		return result, err
	}

	if result.ResponseCode == "400" {
		return coordinate{}, errors.New("city not found")
	}

	return result, nil
}
