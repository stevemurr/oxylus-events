package handler

import (
	"net/http"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

// CreateUser creates a global uuid that keys all their events
func (h *Handler) CreateUser(c echo.Context) error {
	newUser := uuid.NewV4().String()
	h.Users = append(h.Users, newUser)
	return c.JSONPretty(http.StatusCreated, newUser, "  ")
}

// GetUsers returns all users
func (h *Handler) GetUsers(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, h.Users, "  ")
}
