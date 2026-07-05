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
	handler := NewHandler(service)
	api := e.Group("/api/v1/reservations")
	// api.GET("/user", userHandler.GetAllUser)
	api.POST("", handler.ReservationCreate, middlewares.AuthValidation(jwtService), middlewares.Authorized("admin","driver"))

}
