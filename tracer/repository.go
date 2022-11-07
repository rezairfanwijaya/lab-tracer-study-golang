package tracer

import "gorm.io/gorm"

// interface for contract
type TracerRespository interface {
	GetAll() ([]Tracer, error)
	Create(tracer Tracer) error
}

type tracerRespository struct {
	db *gorm.DB
}

func NewTracerRepository(db *gorm.DB) *tracerRespository {
	return &tracerRespository{db}
}

func (r *tracerRespository) GetAll() ([]Tracer, error) {
	var tracers []Tracer
	if err := r.db.Find(&tracers).Error; err != nil {
		return tracers, err
	}

	return tracers, nil
}

func (r *tracerRespository) Create(tracer Tracer) error {
	if err := r.db.Create(&tracer).Error; err != nil {
		return err
	}

	return nil
}
