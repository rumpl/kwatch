package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server is the server
type Server struct {
	e *echo.Echo
}

// NewServer creates a new server
func NewServer() *Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &Server{
		e: e,
	}
}

// Start starts the server
func (s *Server) Start(addr string) {
	s.e.Start(addr)
}
