package user

import (
	"sportsync/internal/auth"
	"sportsync/internal/config"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(e *echo.Echo, db *gorm.DB, env *config.Config) {
	jwtService := auth.NewJwtService(env.JwtSecret, 0)
	userRepo := NewRepository(db)
	userService := NewService(userRepo, jwtService)
	userHandler := NewHandler(userService)
	api := e.Group("/api/v1/auth")
	// api.GET("/user", userHandler.GetAllUser)
	api.POST("/register", userHandler.CreateUser)
	api.POST("/login", userHandler.LoginUser)
}
