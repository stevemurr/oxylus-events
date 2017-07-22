package poller

import (
	"oxylus/driver/driver"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Poller represents an object that repeats a task every n seconds
// This represents a contract with a hardware target
type Poller struct {
	UUID         uuid.UUID     `json:"uuid"`
	Action       string        `json:"action"`
	Timer        *time.Timer   `json:"-"`
	PollInterval time.Duration `json:"pollInterval"`
	Driver       driver.Driver `json:"driver"`
}

// Poll is the loop that polls the hardware
func (p *Poller) Poll() (interface{}, error) {
	response, err := p.Driver.Get(p.Action)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// SetInterval will allow you to set the poll interval
func (p *Poller) SetInterval(val time.Duration) {
	p.PollInterval = val
}

// GetInterval will allow you to set the poll interval
func (p *Poller) GetInterval() time.Duration {
	return p.PollInterval
}

// New returns a new poller
// Pass a reference to the driver you want to poll and it will push the response down the channel C
// The user still needs to call poll before polling begins
func New() *Poller {
	p := &Poller{}
	p.UUID = uuid.NewV4()
	return p
}
