package routes

import (
	"rentalMobil/internal/handlers"
	"rentalMobil/internal/repositories"
	"rentalMobil/internal/services"
	"rentalMobil/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookingRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repositories.NewRepoBooking(db)
	repoCar := repositories.NewRepoCars(db)
	serv := services.NewServiceBooking(repo, repoCar)
	handler := handlers.NewHandlerBooking(serv)

	bookingRoute := r.Group("/bookings", middleware.AuthJWT())
	{
		bookingRoute.POST("", handler.CreateBooking)
		bookingRoute.GET("", handler.GetBookings)
		bookingRoute.GET(":id", handler.GetBooking)
		bookingRoute.PUT(":id", middleware.MustAdmin(), handler.UpdateBooking)
	}
}
