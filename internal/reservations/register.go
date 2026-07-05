package reservations

import (
	"sportsync/internal/auth"
	"sportsync/internal/config"
	"sportsync/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm" 
)

func ReservationRoute(e *echo.Echo, db *gorm.DB, env *config.Config) {
	jwtService := auth.NewJwtService(env.JwtSecret, 0)
	repo := NewRepository(db)
	service := NewService(repo)
	serviceHandler := NewHandler(service)
	api := e.Group("/api/v1/reservations")

	api.GET("", serviceHandler.GetAllReservation, middlewares.AuthValidation(jwtService), middlewares.Authorized("admin"))
	api.POST("", serviceHandler.ReservationCreate, middlewares.AuthValidation(jwtService), middlewares.Authorized("admin", "driver"))
	api.GET("/my-reservations", serviceHandler.MyReservation, middlewares.AuthValidation(jwtService))
	api.DELETE("/:id", serviceHandler.ReservationCancel, middlewares.AuthValidation(jwtService))

}
