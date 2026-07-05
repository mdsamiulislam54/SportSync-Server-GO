package parkingzones

import (
	"sportsync/internal/auth"
	"sportsync/internal/config"
	"sportsync/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func ZoneRoute(e *echo.Echo, db *gorm.DB, env *config.Config) {
	jwtService := auth.NewJwtService(env.JwtSecret, 0)
	repository := NewZoneRepository(db)
	service := NewService(repository)
	zoneHandler := NewHandler(service)

	api := e.Group("/api/v1/zones")
	api.POST("", zoneHandler.CreateZones, middlewares.AuthValidation(jwtService), middlewares.Authorized("admin"))
	api.GET("", zoneHandler.GetAllZone)
	api.GET("/:id", zoneHandler.GetZoneById)

}
