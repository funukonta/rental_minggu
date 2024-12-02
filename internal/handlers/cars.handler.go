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

func (h *HandlerCars) GetCars(c *gin.Context) {
	resp, err := h.serv.GetCars()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *HandlerCars) GetCar(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.serv.GetCar(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *HandlerCars) UpdateCar(c *gin.Context) {
	id := c.Param("id")

	req := new(dtos.UpdateCarReq)
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.serv.UpdateCar(req, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *HandlerCars) DeleteCar(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.serv.DeleteCar(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
