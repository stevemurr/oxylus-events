package handler

import (
	"net/http"
	"oxylus/driver/particleio"

	"github.com/labstack/echo"
)

// GetCommands returns commands
func (h *Handler) GetCommands(c echo.Context) error {
	driver := c.Param("driver")
	switch driver {
	case "particleio":
		d := particleio.New()
		return c.JSONPretty(http.StatusOK, d.Commands(), "  ")
	default:
		return c.NoContent(http.StatusOK)
	}
}
