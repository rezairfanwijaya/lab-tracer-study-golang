package tracer

import (
	"github.com/rezairfanwijaya/lab-tracer-study-golang/coordinat"
)

type TracerService interface {
	ShowAll() ([]Tracer, error)
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

func (s *tracerService) ShowAll() ([]Tracer, error) {
	tracers, err := s.tracerRepo.GetAll()
	if err != nil {
		return tracers, err
	}

	return tracers, nil
}

func (s *tracerService) Save(tracer TracerInput) error {
	// bindcoordinat
	var tracerBind Tracer
	tracerBind.City = tracer.City
	tracerBind.Job = tracer.Job
	tracerBind.Name = tracer.Name

	// get coordinate
	coordinate, err := s.coordinateService.GetCoordinate(tracerBind.City)
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
