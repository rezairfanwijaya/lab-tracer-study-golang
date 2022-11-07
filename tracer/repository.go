package tracer

import "gorm.io/gorm"

// interface for contract
type TracerRespository interface {
	GetAll() ([]Tracer, error)
}

type tracerRespository struct {
	db *gorm.DB
}

func NewTracerRepository(db *gorm.DB) *tracerRespository {
	return &tracerRespository{db}
}

func (t *tracerRespository) GetAll() ([]Tracer, error) {
	var tracers []Tracer
	if err := t.db.Find(&tracers).Error; err != nil {
		return tracers, err
	}

	return tracers, nil
}
