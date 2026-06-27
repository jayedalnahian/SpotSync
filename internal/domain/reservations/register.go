package reservations

import (
	"SpotSync/internal/auth"
	"SpotSync/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, jwtService auth.JWTService) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	api := e.Group("/api/v1/reservations")

	api.POST("", handler.ReserveSpot, middlewares.AuthMiddleware(jwtService))
}
