package handler

import (
	"net/http"
	"oxylus/eventregistry"
	"oxylus/pollerregistry"
	"oxylus/store/boltstore"

	mgo "gopkg.in/mgo.v2"

	"github.com/labstack/echo"
)

type (
	// Handler is our global state to the event registry and database
	Handler struct {
		PollerRegistry *pollerregistry.PollerRegistry
		EventRegistry  *eventregistry.EventRegistry
		Store          *boltstore.BoltStore
		DB             *mgo.Session
		Users          []string
	}
)

// Test tests the handlers
func (h *Handler) Test(c echo.Context) error {
	return c.JSON(http.StatusOK, "Wide Diaper")
}
