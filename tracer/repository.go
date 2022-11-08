package tracer

import "gorm.io/gorm"

// interface for contract
type TracerRespository interface {
	GetAll() ([]Tracer, error)
	Create(tracer Tracer) error
	GetTracerBaseCity() ([]TracerBaseCity, error)
}

type tracerRespository struct {
	db *gorm.DB
}

func NewTracerRepository(db *gorm.DB) *tracerRespository {
	return &tracerRespository{db}
}

// show all tracer
func (r *tracerRespository) GetAll() ([]Tracer, error) {
	var tracers []Tracer
	if err := r.db.Find(&tracers).Error; err != nil {
		return tracers, err
	}

	return tracers, nil
}

// save new tracer
func (r *tracerRespository) Create(tracer Tracer) error {
	if err := r.db.Create(&tracer).Error; err != nil {
		return err
	}

	return nil
}

// show tracer group by city
func (r *tracerRespository) GetTracerBaseCity() ([]TracerBaseCity, error) {
	var tracerBaseCity []TracerBaseCity

	sql := `SELECT id, city, latitude, longitude, count(city) as total from tracers GROUP by city`
	r.db.Raw(sql).Scan(&tracerBaseCity)

	return tracerBaseCity, nil
}
