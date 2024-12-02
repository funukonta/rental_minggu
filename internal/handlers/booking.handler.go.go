package handlers

import (
	"net/http"
	"rentalMobil/internal/dtos"
	"rentalMobil/internal/services"

	"github.com/gin-gonic/gin"
)

type HandlerBooking struct {
	serv *services.ServiceBooking
}

func NewHandlerBooking(serv *services.ServiceBooking) *HandlerBooking {
	return &HandlerBooking{serv: serv}
}

func (h *HandlerBooking) CreateBooking(c *gin.Context) {
	req := new(dtos.CreateBookingReq)
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.serv.CreateBooking(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
