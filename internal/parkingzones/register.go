package parkingzones

import (
	"sportsync/internal/config"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func ZoneRoute(e *echo.Echo, db *gorm.DB, env *config.Config) {

	repository := NewZoneRepository(db)
	service := NewService(repository)
	zoneHandler := NewHandler(service)

	api := e.Group("/api/v1/zones")
	api.POST("", zoneHandler.CreateZones)
	api.GET("", zoneHandler.GetAllZone)
	api.GET("/:id", zoneHandler.GetZoneById)

}
