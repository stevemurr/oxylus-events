package handler

import (
	"net/http"
	"oxylus/eventregistry"
	"oxylus/pollerregistry"
	"oxylus/store/boltstore"

	"github.com/labstack/echo"
)

type (
	// Handler is our global state to the event registry and database
	Handler struct {
		PollerRegistry *pollerregistry.PollerRegistry
		EventRegistry  *eventregistry.EventRegistry
		Store          *boltstore.BoltStore
		Users          []string
	}
)

// Test tests the handlers
func (h *Handler) Test(c echo.Context) error {
	return c.JSON(http.StatusOK, "Wide Diaper")
}
