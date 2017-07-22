package eventregistry

import (
	"errors"
	"log"
	"oxylus/event"
	"sync"
	"time"
)

var (
	// ErrKeyNotFound denotes the key does not exist in the registry
	ErrKeyNotFound = errors.New("key not found")
	// ErrElementNotFound means the event id does not exist
	ErrElementNotFound = errors.New("element not found in registry")
	// ErrTimeIntervalLessThanZero means the time cant be scheduled
	ErrTimeIntervalLessThanZero = errors.New("time intveral is less than zero")
)

// EventRegistry uses a uuid to key a list of events
type EventRegistry struct {
	Registry map[string]map[string]*event.Event
	sync.Mutex

	TimerStarted chan string
	TimerEnded   chan string
}

// New returns a new registry
func New() *EventRegistry {
	return &EventRegistry{
		Registry:     make(map[string]map[string]*event.Event),
		TimerStarted: make(chan string),
		TimerEnded:   make(chan string),
	}
}

// StartTimer allows us to place middleware around the event
// This is useful for sending info to channels on completion etc.
func (e *EventRegistry) StartTimer(key, itemID string) error {
	event, err := e.Get(key, itemID)
	if err != nil {
		return ErrElementNotFound
	}
	if event.TimeInterval <= 0 {
		return ErrTimeIntervalLessThanZero
	}
	event.Status = "timer started"
	e.TimerStarted <- event.String()
	event.Timer = time.AfterFunc(event.TimeInterval, func() {
		event.Status = "timer ended"
		e.TimerEnded <- event.String()
		if err := event.Driver.Run(event.Action); err != nil {
			log.Println(err)
		}
		if event.Repeats {
			e.StartTimer(key, itemID)
		}
	})
	return nil
}

// StopTimer allows us to place middleware around the event
// This is useful for pruning a dead event
func (e *EventRegistry) StopTimer(key, itemID string) error {
	event, err := e.Get(key, itemID)
	if err != nil {
		return ErrElementNotFound
	}
	event.Status = "timer stopped"
	event.StopTimer()
	e.TimerEnded <- event.String()
	return nil
}

// GetAll returns the map against the user uuid
func (e *EventRegistry) GetAll(key string) map[string]*event.Event {
	if val, ok := e.Registry[key]; ok {
		return val
	}
	return map[string]*event.Event{}
}

// Get returns a single event
func (e *EventRegistry) Get(key, itemID string) (*event.Event, error) {
	if val, ok := e.Registry[key][itemID]; ok {
		return val, nil
	}
	return nil, ErrElementNotFound
}

// Add sets a value in the registry
func (e *EventRegistry) Add(key string, val *event.Event) {
	e.Lock()
	if _, ok := e.Registry[key]; ok {
		e.Registry[key][val.UUID.String()] = val
	} else {
		e.Registry[key] = map[string]*event.Event{
			val.UUID.String(): val,
		}
	}
	e.Unlock()
}

// RemoveEvent will remove an event from the map
func (e *EventRegistry) RemoveEvent(key, itemID string) error {
	if _, ok := e.Registry[key][itemID]; !ok {
		e.Lock()
		delete(e.Registry[key], itemID)
		e.Unlock()
	}
	return nil
}
