package routes

import (
	"rentalMobil/internal/handlers"
	"rentalMobil/internal/repositories"
	"rentalMobil/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookingRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repositories.NewRepoBooking(db)
	repoCar := repositories.NewRepoCars(db)
	serv := services.NewServiceBooking(repo, repoCar)
	handler := handlers.NewHandlerBooking(serv)

	bookingRoute := r.Group("/booking")
	{
		bookingRoute.POST("", handler.CreateBooking)

	}
}
