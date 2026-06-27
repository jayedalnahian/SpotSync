package reservations

import (
	"SpotSync/internal/domain/reservations/dto"
	"SpotSync/internal/httpresponse"
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) ReserveSpot(c *echo.Context) error {
	// Authentication verification
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized: User ID missing",
		})
	}

	var req dto.ReserveSpotRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	res, err := h.service.ReserveSpot(userID, req)
	if err != nil {
		if errors.Is(err, ErrZoneFull) {
			return c.JSON(http.StatusConflict, httpresponse.Error{
				Code:    http.StatusConflict,
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to reserve spot",
			Details: err.Error(),
		})
	}

	result := httpresponse.SendResponse{
		Success: true,
		Message: "Reservation confirmed successfully",
		Data:    res,
	}
	return c.JSON(http.StatusCreated, result)
}

func (h *handler) GetMyReservations(c *echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized: User ID missing",
		})
	}

	res, err := h.service.GetMyReservations(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve reservations",
			Details: err.Error(),
		})
	}

	result := httpresponse.SendResponse{
		Success: true,
		Message: "My reservations retrieved successfully",
		Data:    res,
	}
	return c.JSON(http.StatusOK, result)
}