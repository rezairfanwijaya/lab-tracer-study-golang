package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezairfanwijaya/lab-tracer-study-golang/helper"
	"github.com/rezairfanwijaya/lab-tracer-study-golang/tracer"
)

type tracerHandler struct {
	tracerService tracer.TracerService
}

func NewTracerHandler(tracerService tracer.TracerService) *tracerHandler {
	return &tracerHandler{tracerService}
}

// handler to show all tracer data
func (h *tracerHandler) GetAll(c *gin.Context) {
	// call service
	tracers, err := h.tracerService.ShowAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, tracers)
		return
	}

	// return response
	response := tracer.TracerFormater(tracers)
	c.JSON(http.StatusOK, response)
}

// handler to save new tracer data
func (h *tracerHandler) SaveTracer(c *gin.Context) {
	// binding
	var inputTracer tracer.TracerInput

	err := c.ShouldBindJSON(&inputTracer)
	if err != nil {
		myError := helper.ErrorBinding(err)
		response := helper.ResponseAPI("failed", http.StatusBadRequest, myError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// call service
	if err := h.tracerService.Save(inputTracer); err != nil {
		response := helper.ResponseAPI("failed", http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// return response
	response := helper.ResponseAPI("success", http.StatusOK, "success")
	c.JSON(http.StatusCreated, response)
}

// handler to show tracer data group by city
func (h *tracerHandler) GetAllTracerBaseCity(c *gin.Context) {
	// call service
	tracerBaseCity, err := h.tracerService.ShowBaseCity()
	if err != nil {
		response := helper.ResponseAPI("failed", http.StatusInternalServerError, err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// return response
	response := helper.ResponseAPI("success", http.StatusOK, tracerBaseCity)
	c.JSON(http.StatusOK, response)
}
