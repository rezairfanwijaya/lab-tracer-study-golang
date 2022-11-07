package tracer

type TracerInput struct {
	Name string `json:"name" binding:"required"`
	City string `json:"city" binding:"required"`
	Job  string `json:"job" binding:"required"`
}
