package middlewares

import (
	"net/http"
	"sportsync/internal/auth"
	"strings"

	"github.com/labstack/echo/v5"
)

func AuthValidation(jwtService auth.JwtService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Missing authorization header",
				})
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Missing authorization bearer token",
				})
			}

			token := parts[1]

			claims, err := jwtService.TokenValidation(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid or expire token",
				})
			}

			c.Set("userId", claims.UserID)
			c.Set("name", claims.Name)
			c.Set("email", claims.Email)

			return next(c)

		}
	}
}
