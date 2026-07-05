package reservations

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

func (h handler) ReservationCreate(c *echo.Context) error {
	var req Reservation
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to bind request",
			Details: err.Error(),
		})
	}

	userID, ok := c.Get("userId").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	reservation := &Reservation{
		UserID:       userID,
		ZoneID:       req.ZoneID,
		LicensePlate: req.LicensePlate,
	}

	res, err := h.service.ReservationCreate(reservation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create Reservations",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, res)

}

func (h handler) MyReservation(c *echo.Context) error {
	userId, ok := c.Get("userId").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	res, err := h.service.MyReservation(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to Retrieved Reservation",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)

}

func (h handler) ReservationCancel(c *echo.Context) error {
	role, ok := c.Get("role").(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "user role not found")
	}

	userID, ok := c.Get("userId").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "user id not found")
	}

	reservationID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid reservation id")
	}

	if err := h.service.CancelReservation(uint(reservationID), uint(userID), role); err != nil {
		return err
	}

	return c.JSON(http.StatusInternalServerError, map[string]any{
		"success": true,
		"message": "Reservation cancelled successfully",
	})

}
func (h handler) GetAllReservation(c *echo.Context) error {
	role, ok := c.Get("role").(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "user role not found")
	}

	response, err := h.service.GetAllReservation(role)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"success": false,
			"message": "Reservation data not found",
			"errors":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"data":    response,
		"success": true,
		"message": "Reservation Retrieved successfully",
	})

}
