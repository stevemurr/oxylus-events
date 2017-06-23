package poller

import (
	"oxylus/driver/driver"
	"time"
)

// Poller represents an object that repeats a task every n seconds
// This represents a contract with a hardware target
type Poller struct {
	Driver       driver.Driver
	Action       string `json:"action"`
	timer        *time.Timer
	C            chan interface{}
	PollInterval time.Duration `json:"pollInterval"`
}

// Poll is the loop that polls the hardware
func (p *Poller) Poll(val interface{}) {
	p.timer = time.AfterFunc(p.PollInterval, func() {
		if err := p.Driver.Get(p.Action, val); err == nil {
			p.C <- val
		}
		p.Poll(val)
	})
}

// SetInterval will allow you to set the poll interval
func (p *Poller) SetInterval(val time.Duration) {
	p.PollInterval = val
}

// GetInterval will allow you to set the poll interval
func (p *Poller) GetInterval() time.Duration {
	return p.PollInterval
}

// NewPoller returns a new poller
// Pass a reference to the driver you want to poll and it will push the response down the channel C
// The user still needs to call poll before polling begins
func NewPoller(action string, driver driver.Driver, c chan interface{}, pollInterval time.Duration) *Poller {
	p := &Poller{}
	p.Action = action
	p.Driver = driver
	p.C = c
	p.PollInterval = pollInterval
	return p
}
