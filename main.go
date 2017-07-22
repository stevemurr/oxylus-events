package main

import (
	"log"
	"oxylus/eventregistry"
	"oxylus/handler"
	"oxylus/pollerregistry"
	"oxylus/store/boltstore"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"

	"oxylus/driver/particleio"

	mw "github.com/labstack/echo/middleware"
)

// func hello(c echo.Context) error {
// 	websocket.Handler(func(ws *websocket.Conn) {
// 		defer ws.Close()
// 		for {
// 			// Write
// 			err := websocket.Message.Send(ws, "Hello, Client!")
// 			if err != nil {
// 				c.Logger().Error(err)
// 			}

// 			// Read
// 			msg := ""
// 			err = websocket.Message.Receive(ws, &msg)
// 			if err != nil {
// 				c.Logger().Error(err)
// 			}
// 			fmt.Printf("%s\n", msg)
// 		}
// 	}).ServeHTTP(c.Response(), c.Request())
// 	return nil
// }

func main() {
	e := echo.New()

	e.Use(mw.Logger())
	e.Use(mw.CORS())
	e.Use(mw.Recover())
	b, err := boltstore.NewStore("bolt.db")
	if err != nil {
		log.Fatal(err)
	}

	h := handler.Handler{
		Store:          b,
		EventRegistry:  eventregistry.New(),
		PollerRegistry: pollerregistry.New(),
		Users:          []string{},
	}

	e.GET("/", h.Test)

	e.GET("/events", h.GetAllEvents)

	e.GET("/user", h.GetUsers)
	e.POST("/user", h.CreateUser)

	e.GET("/user/:id/events", h.GetUserEvents)
	e.GET("/user/:id/events/:event", h.GetUserEvent)
	e.POST("/user/:id/events", h.CreateEvent)
	e.DELETE("/user/:id/events/:event", h.DeleteEvent)

	e.GET("/search", h.Find)

	// TODO: polling handlers
	e.GET("/user/:id/poller", h.GetPollers)
	e.POST("/user/:id/poller", h.CreatePoller)
	e.GET("/user/:id/poller/:poller", h.GetPoller)
	e.DELETE("/user/:id/poller/:poller", h.DeletePoller)

	// TODO: driver handlers
	e.GET("/drivers/:driver/commands", h.GetCommands)
	// e.GET("/ws", hello)
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

	go func(h *handler.Handler) {
		for {
			select {
			case msg := <-h.PollerRegistry.ToDB:
				r := msg.(*particleio.OxylusResponse)
				log.Println("[POLLED] " + r.String())
				h.Store.Upsert(uuid.NewV4().String(), r)
			}
		}
	}(&h)

	e.Logger.Fatal(e.Start(":1323"))
}
