package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func Authorized(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			role, ok := c.Get("role").(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Unauthorized access",
				})
			}

			allowed := map[string]struct{}{
				"admin":  {},
				"driver": {},
				"user":   {},
			}

			if _, ok := allowed[role]; !ok {
				return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
			}

			return next(c)
			
		}
	}
}
