package handler

import (
	"fmt"
	"net/http"

	"oxylus/driver/particleio"

	"github.com/labstack/echo"
)

// TODO: Implement a normalized response object rather than driver specific ones
// TODO: Add type or DB switch

// Find lets you search the database for a field and value
// You will generally be search the database for response data to graph etc.
func (h *Handler) Find(c echo.Context) error {
	field := c.QueryParam("field")
	value := c.QueryParam("value")
	var results []particleio.Response
	if err := h.Store.Find(field, value, &results); err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSONPretty(http.StatusOK, results, "  ")
}
