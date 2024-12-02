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

func (h *HandlerBooking) GetBookings(c *gin.Context) {
	userid, exists := c.Get("userid") // check tipe data userid = any
	if !exists {
		userid = "" // kalo ngga ada force to empty string
	}
	resp, err := h.serv.GetBookings(userid.(string)) // kenapa pake .(string) -> ini merubah userid dengan tipedata any ke tipe data string (menegaskan any itu string!)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *HandlerBooking) GetBooking(c *gin.Context) {
	userid, exists := c.Get("userid") // check tipe data userid = any
	if !exists {
		userid = "" // kalo ngga ada force to empty string
	}
	id := c.Param("id")

	resp, err := h.serv.GetBooking(userid.(string), id) // kenapa pake .(string) -> ini merubah userid dengan tipedata any ke tipe data string (menegaskan any itu string!)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *HandlerBooking) UpdateBooking(c *gin.Context) {
	userid, exists := c.Get("userid") // check tipe data userid = any
	if !exists {
		userid = "" // kalo ngga ada force to empty string
	}
	id := c.Param("id")
	req := &dtos.UpdateBookingReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.serv.UpdateBooking(userid.(string), id, req) // kenapa pake .(string) -> ini merubah userid dengan tipedata any ke tipe data string (menegaskan any itu string!)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
