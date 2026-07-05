package parkingzones

import (
	"net/http"
	"sportsync/internal/httpresponse"
	"strconv"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{
		service,
	}
}

func (h handler) CreateZones(c *echo.Context) error {

	var req Zone
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to bind request",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to validate request",
			Details: err.Error(),
		})
	}

	res, err := h.service.CreateZones(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create zone",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, res)

}

func (h handler) GetAllZone(c *echo.Context) error {

	res, err := h.service.GetAllZones()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to Retrieved zone",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)

}
func (h handler) GetZoneById(c *echo.Context) error {
	zoneId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid zone id",
			Details: err.Error(),
		})
	}
	res, err := h.service.GetZoneById(uint(zoneId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to Retrieved zone",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)

}
