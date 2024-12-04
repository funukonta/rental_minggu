package routes

import (
	"rentalMobil/internal/handlers"
	"rentalMobil/internal/repositories"
	"rentalMobil/internal/services"
	"rentalMobil/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CarsRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repositories.NewRepoCars(db)
	serv := services.NewServiceCars(repo)
	handler := handlers.NewHandlerCars(serv)

	carsRoute := r.Group("/cars", middleware.AuthJWT())
	{
		carsRoute.POST("", middleware.MustAdmin(), handler.CreateCar)
		carsRoute.GET(":id", handler.GetCar) // localhost:8080/cars/id
		carsRoute.GET("", handler.GetCars)

		carsRoute.PUT(":id", middleware.MustAdmin(), handler.UpdateCar)
		carsRoute.DELETE(":id", middleware.MustAdmin(), handler.DeleteCar)

	}
}
