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

func (h *tracerHandler) GetAll(c *gin.Context) {
	tracers, err := h.tracerService.ShowAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, tracers)
		return
	}

	response := tracer.TracerFormater(tracers)
	c.JSON(http.StatusOK, response)
}

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

	if err := h.tracerService.Save(inputTracer); err != nil {
		response := helper.ResponseAPI("failed", http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseAPI("success", http.StatusOK, "success")
	c.JSON(http.StatusOK, response)
}
