package parking_zones

import (
	"SpotSync/internal/httpresponse"
	"SpotSync/internal/domain/parking_zones/dto"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service service) *handler {
	return &handler{
		service: &service,
	}
}

func (h *handler) CreateZone(c *echo.Context) error {
	// Role verification
	role, ok := c.Get("user_role").(string)
	if !ok || role != "admin" {
		return c.JSON(http.StatusForbidden, httpresponse.Error{
			Code:    http.StatusForbidden,
			Message: "Forbidden: Admin access required",
		})
	}

	var req dto.CreateZoneRequest
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

	res, err := h.service.CreateZone(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create parking zone",
			Details: err.Error(),
		})
	}

	result := httpresponse.SendResponse{
		Success: true,
		Message: "Parking zone created successfully",
		Data:    res,
	}
	return c.JSON(http.StatusCreated, result)
}
