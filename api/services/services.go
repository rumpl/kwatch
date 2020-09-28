package services

import (
	"github.com/labstack/echo/v4"
)

// Services is the api for the services
type Services struct {
}

// Service represents a running service
type Service struct {
	Name  string
	Image string
}

// New creates a new services thing
func New() *Service {
	return &Service{}
}

// Register adds all the services routes to the router
func (s *Service) Register(router *echo.Group) {
	router.Add("GET", "/services", s.List)
}

// List returns all the running services
func (s *Service) List(ctx echo.Context) error {
	return nil
}
