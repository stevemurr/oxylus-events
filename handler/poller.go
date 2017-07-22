package handler

import (
	"log"
	"net/http"
	"oxylus/poller"
	"time"

	"github.com/labstack/echo"
)

type (
	// PollerRequest --
	PollerRequest struct {
		Action       string            `json:"action"`
		Driver       string            `json:"driver"`
		DriverParams map[string]string `json:"driverParams"`
		PollInterval string            `json:"pollInterval"`
	}
	// PollerResponse --
	PollerResponse struct {
		Poller  *poller.Poller `json:"poller"`
		Request *PollerRequest `json:"request"`
	}
)

// CreatePoller creates a poller
func (h *Handler) CreatePoller(c echo.Context) error {
	id := c.Param("id")
	request := new(PollerRequest)
	var err error
	if err = c.Bind(request); err != nil {
		return err
	}
	p := poller.New()
	p.Action = request.Action
	p.Driver = NewDriver(request.Driver, request.DriverParams)
	p.PollInterval, err = time.ParseDuration(request.PollInterval)
	if err != nil {
		return err
	}
	response := new(PollerResponse)
	h.PollerRegistry.Add(id, p)
	if err = h.PollerRegistry.Poll(id, p.UUID.String()); err != nil {
		log.Println(err)
	}
	response.Poller = p
	response.Request = request
	return c.JSONPretty(http.StatusCreated, response, "  ")
}

// GetPoller --
func (h *Handler) GetPoller(c echo.Context) error {
	id := c.Param("id")
	poller := c.Param("poller")
	p, err := h.PollerRegistry.Get(id, poller)
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, p, "  ")
}

// DeletePoller --
func (h *Handler) DeletePoller(c echo.Context) error {
	id := c.Param("id")
	poller := c.Param("poller")
	if err := h.PollerRegistry.RemovePoller(id, poller); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

// GetPollers --
func (h *Handler) GetPollers(c echo.Context) error {
	id := c.Param("id")
	return c.JSONPretty(http.StatusOK, h.PollerRegistry.Registry[id], "  ")
}
