package server

import (
	"SpotSync/internal/auth"
	"SpotSync/internal/config"
	"SpotSync/internal/domain/parking_zones"
	"fmt"
	"net/http"
	"SpotSync/internal/domain/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func Start(db *gorm.DB, cfg *config.Config) {
	db.AutoMigrate(&user.User{}, &parking_zones.ParkingZone{})

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.RequestLogger())

	e.GET("/health", func(c *echo.Context) error {
		return c.String(http.StatusOK, "running")
	})

	jwtService := auth.NewJWTService(cfg.JwtSecret)

	//routes
	user.RegisterRoutes(e, db, cfg)
	parking_zones.RegisterRoutes(e, db, jwtService)

	port := fmt.Sprintf(":%s", cfg.Port)
	if err := e.Start(port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}