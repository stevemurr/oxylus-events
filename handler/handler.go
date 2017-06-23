package handler

import (
	"net/http"
	"oxylus/eventregistry"
	"oxylus/store/boltstore"

	"github.com/labstack/echo"
)

type (
	// Handler is our global state to the event registry and database
	Handler struct {
		EventRegistry *eventregistry.EventRegistry
		Store         *boltstore.BoltStore
	}
)

// Test tests the handlers
func (h *Handler) Test(c echo.Context) error {
	return c.JSON(http.StatusOK, "Wide Diaper")
}
