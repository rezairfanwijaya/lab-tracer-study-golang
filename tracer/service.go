package tracer

import (
	"strings"

	"github.com/rezairfanwijaya/lab-tracer-study-golang/coordinat"
)

type TracerService interface {
	ShowAll() ([]Tracer, error)
	ShowBaseCity() ([]TracerBaseCity, error)
	Save(tracer TracerInput) error
}

type tracerService struct {
	tracerRepo        TracerRespository
	coordinateService coordinat.CoordinateService
}

func NewTracerService(
	tracerRepo TracerRespository,
	coordinateService coordinat.CoordinateService,
) *tracerService {
	return &tracerService{tracerRepo, coordinateService}
}

// show all tracer data
func (s *tracerService) ShowAll() ([]Tracer, error) {
	tracers, err := s.tracerRepo.GetAll()
	if err != nil {
		return tracers, err
	}

	return tracers, nil
}

// save new tracer data
func (s *tracerService) Save(tracer TracerInput) error {
	// bindcoordinat
	var tracerBind Tracer
	tracerBind.City = tracer.City
	tracerBind.Job = tracer.Job
	tracerBind.Name = tracer.Name

	city := strings.ReplaceAll(tracerBind.City, " ", "")

	// get coordinate
	coordinate, err := s.coordinateService.GetCoordinate(city)
	if err != nil {
		return err
	}

	// assign coordinate
	tracerBind.Latitude = coordinate.Latitude
	tracerBind.Longitude = coordinate.Longitude

	if err := s.tracerRepo.Create(tracerBind); err != nil {
		return err
	}

	return nil
}

// show tracer data group by city
func (s *tracerService) ShowBaseCity() ([]TracerBaseCity, error) {
	tracerBaseCity, err := s.tracerRepo.GetTracerBaseCity()
	if err != nil {
		return tracerBaseCity, err
	}

	return tracerBaseCity, nil
}
