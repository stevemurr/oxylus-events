package main

import (
	"log"
	"oxylus/eventregistry"
	"oxylus/handler"
	"oxylus/store/boltstore"

	"github.com/labstack/echo"

	mw "github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(mw.Logger())
	e.Use(mw.CORS())
	e.Use(mw.Recover())
	b, err := boltstore.NewStore("bolt.db")
	if err != nil {
		log.Fatal(err)
	}

	h := handler.Handler{Store: b, EventRegistry: eventregistry.New()}

	e.GET("/", h.Test)

	e.GET("/user/:id", h.GetUserEvents)
	e.POST("/user", h.CreateUser)

	e.POST("/user/:id/event", h.CreateEvent)
	e.DELETE("/user/:id/event/:event", h.DeleteEvent)

	go func(r *eventregistry.EventRegistry) {
		for {
			select {
			case msg := <-r.TimerStarted:
				log.Println("[STARTING] " + msg)
			case msg := <-r.TimerEnded:
				log.Println("[ENDING]   " + msg)
			}
		}
	}(h.EventRegistry)

	e.Logger.Fatal(e.Start(":1323"))
}
