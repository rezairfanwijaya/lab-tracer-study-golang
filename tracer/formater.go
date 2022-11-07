package tracer

type tracerFormater struct {
	Name string `json:"name"`
	City string `json:"city"`
	Job  string `json:"job"`
}

func singleTracerFormater(tracer Tracer) tracerFormater {
	var tracerFormated tracerFormater
	tracerFormated.City = tracer.City
	tracerFormated.Job = tracer.Job
	tracerFormated.Name = tracer.Name

	return tracerFormated
}

func TracerFormater(tracers []Tracer) []tracerFormater {
	var result []tracerFormater
	for _, tracer := range tracers {
		singleTracerFormated := singleTracerFormater(tracer)
		result = append(result, singleTracerFormated)
	}

	return result
}
