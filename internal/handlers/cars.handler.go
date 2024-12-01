package handlers

import (
	"net/http"
	"rentalMobil/internal/dtos"
	"rentalMobil/internal/services"

	"github.com/gin-gonic/gin"
)

type HandlerCars struct {
	serv *services.ServiceCars
}

func NewHandlerCars(serv *services.ServiceCars) *HandlerCars {
	return &HandlerCars{serv: serv}
}

func (h *HandlerCars) CreateCar(c *gin.Context) {
	req := new(dtos.CreateCarReq)
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Normalize()

	resp, err := h.serv.CreateCars(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
