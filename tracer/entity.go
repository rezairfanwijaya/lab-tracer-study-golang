package tracer

type Tracer struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	Job  string `json:"job"`
}
