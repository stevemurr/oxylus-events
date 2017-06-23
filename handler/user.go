package handler

import (
	"net/http"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

// GetUserEvents returns all of a users events
func (h *Handler) GetUserEvents(c echo.Context) error {
	id := c.Param("id")
	return c.JSONPretty(http.StatusOK, h.EventRegistry.GetAll(id), "  ")
}

// CreateUser creates a global uuid that keys all their events
func (h *Handler) CreateUser(c echo.Context) error {
	return c.JSONPretty(http.StatusCreated, uuid.NewV4().String(), "  ")
}
