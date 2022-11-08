package tracer

type TracerBaseCity struct {
	ID        int    `json:"id"`
	City      string `json:"city"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Total     int    `json:"total"`
}
