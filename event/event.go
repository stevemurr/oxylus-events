package event

import (
	"errors"
	"fmt"
	"oxylus/driver/driver"
	"time"

	uuid "github.com/satori/go.uuid"
)

var (
	// ErrTimeIntervalNotPositive means the timer can't activate a function in the future
	ErrTimeIntervalNotPositive = errors.New("time interval is less than zero")
)

// Event represents an action to be executed at some time
type Event struct {
	UUID         uuid.UUID     `json:"uuid"`
	CreatedAt    time.Time     `json:"createdAt"`
	FinishAt     time.Time     `json:"finishAt"`
	Action       string        `json:"action"`
	Driver       driver.Driver `json:"driver"`
	Repeats      bool          `json:"repeats"`
	TimeInterval time.Duration `json:"timeInterval"`
	Timer        *time.Timer   `json:"-"`
}

// StopTimer attempts to stop the time and returns a bool indicating status
func (e *Event) StopTimer() bool {
	return e.Timer.Stop()
}

// StartTimer starts the timer
func (e *Event) StartTimer() error {
	if e.TimeInterval <= 0.0 {
		return ErrTimeIntervalNotPositive
	}
	if err := e.Driver.Run(e.Action); err != nil {
		return err
	}
	return nil
}

func (e *Event) String() string {
	return fmt.Sprintf("%s %s %s %t %s\n", e.UUID.String(), e.Driver.Name(), e.Action, e.Repeats, e.FinishAt)
}

// New returns a new event
func New() *Event {
	return &Event{
		UUID:      uuid.NewV4(),
		CreatedAt: time.Now(),
	}
}
