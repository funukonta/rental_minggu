package handlers

import (
	"net/http"
	"rentalMobil/internal/dtos"
	"rentalMobil/internal/services"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	serv *services.ServiceUser
}

func NewHandlerUser(serv *services.ServiceUser) *HandlerUser {
	return &HandlerUser{serv: serv}
}

func (h *HandlerUser) Login(c *gin.Context) {
	req := new(dtos.LoginReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.serv.Login(*req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *HandlerUser) Register(c *gin.Context) {
	req := new(dtos.RegisterReq)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.serv.Register(*req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
