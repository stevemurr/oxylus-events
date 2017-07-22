package pollerregistry

import (
	"errors"
	"log"
	"oxylus/poller"
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

// PollerRegistry uses a uuid to key a list of events
type PollerRegistry struct {
	Registry map[string]map[string]*poller.Poller
	sync.Mutex
	ToDB           chan interface{}
	PollerStarted  chan string
	PollerExecuted chan string
}

// New returns a new registry
func New() *PollerRegistry {
	return &PollerRegistry{
		Registry:       make(map[string]map[string]*poller.Poller),
		PollerStarted:  make(chan string),
		PollerExecuted: make(chan string),
		ToDB:           make(chan interface{}),
	}
}

// Poll --
func (e *PollerRegistry) Poll(key, itemID string) error {
	poller, err := e.Get(key, itemID)
	if err != nil {
		return ErrElementNotFound
	}
	poller.Timer = time.AfterFunc(poller.GetInterval(), func() {
		response, err := poller.Poll()
		if err != nil {
			log.Println(err)
			// this will not restart the poller
			// use exponential retry timer
		} else {
			e.Poll(key, itemID)
			e.ToDB <- response
		}
	})
	return nil
}

// StopPoller allows us to place middleware around the event
// This is useful for pruning a dead event
func (e *PollerRegistry) StopPoller(key, itemID string) error {
	poller, err := e.Get(key, itemID)
	if err != nil {
		return ErrElementNotFound
	}
	poller.Timer.Stop()
	return nil
}

// GetAll returns the map against the user uuid
func (e *PollerRegistry) GetAll(key string) map[string]*poller.Poller {
	if val, ok := e.Registry[key]; ok {
		return val
	}
	return map[string]*poller.Poller{}
}

// Get returns a single event
func (e *PollerRegistry) Get(key, itemID string) (*poller.Poller, error) {
	if val, ok := e.Registry[key][itemID]; ok {
		return val, nil
	}
	return nil, ErrElementNotFound
}

// Add sets a value in the registry
func (e *PollerRegistry) Add(key string, val *poller.Poller) {
	e.Lock()
	if _, ok := e.Registry[key]; ok {
		e.Registry[key][val.UUID.String()] = val
	} else {
		e.Registry[key] = map[string]*poller.Poller{
			val.UUID.String(): val,
		}
	}
	e.Unlock()
}

// RemovePoller will remove an event from the map
func (e *PollerRegistry) RemovePoller(key, itemID string) error {
	if _, ok := e.Registry[key][itemID]; !ok {
		e.Lock()
		delete(e.Registry[key], itemID)
		e.Unlock()
	}
	return nil
}
