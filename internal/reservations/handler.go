package reservations

import (
	"fmt"
	"net/http"
	"sportsync/internal/httpresponse"

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

	fmt.Println("reservation ...................",reservation)
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
