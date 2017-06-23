package particleio

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

var (
	// ErrActionDoesNotExist means the requested action is not available
	ErrActionDoesNotExist = errors.New("action does not exist")
	// ErrStateChangeRequestFailed means the requested action failed to return 200
	ErrStateChangeRequestFailed = errors.New("particle io service did not return 200")
	// ErrRequestCreationFailed means the request could not be built - probably an error in the request body
	ErrRequestCreationFailed = errors.New("problem creating http request")
	// ErrGetRequestFailed means the request failed
	ErrGetRequestFailed = errors.New("problem sending a get request")
)

// ParticleIO implements driver interface for particleio board
type ParticleIO struct {
	UUID        uuid.UUID `json:"uuid"`
	DeviceID    string    `json:"deviceId"`
	AccessToken string    `json:"accessToken"`
}

// Authenticate hook performs any auth
// Get a new access token at this step
func (p *ParticleIO) Authenticate() error {
	p.DeviceID = "290018000347353137323334"
	p.AccessToken = "f19c897a5232a616fb611500e8fd62951858566a"
	return nil
}

// Raw implements the driver interface
func (p *ParticleIO) Raw(args ...string) error {
	// parse raw args and do what you want!
	fmt.Println(args)
	return nil
}

// Get returns a response to the client with data
// This is used internally to poll the sensors on a regular interval
// The user pulls out data from the store api
func (p *ParticleIO) Get(action string, val interface{}) error {
	if action == "test" {
		fmt.Println("test get method")
		val = Response{}
		return nil
	}
	u := fmt.Sprintf("https://api.particle.io/v1/devices/%s/%s?access_token=%s", p.DeviceID, action, p.AccessToken)

	resp, err := http.Get(u)
	if err != nil {
		return ErrGetRequestFailed
	}
	if err := json.NewDecoder(resp.Body).Decode(val); err != nil {
		return err
	}
	// Inject any data we want into the structure here
	defer resp.Body.Close()
	return nil
}

// Run implements the driver interface and is invoked when the user wants to alter state
// Inversely, call Get when you just want to read sensor data
func (p *ParticleIO) Run(action string) error {
	if action == "test" {
		log.Printf("[RUN] test method for %s\n", p.Name())
		return nil
	}

	u := fmt.Sprintf("https://api.particle.io/v1/products/4439/devices/%s/control", p.DeviceID)
	m := "POST"
	if err := p.makeRequest(u, m, action); err != nil {
		return err
	}
	return nil
}

// Commands implements the driver interface
func (p *ParticleIO) Commands() []map[string]string {
	return []map[string]string{
		{
			"cmd":         "lights",
			"description": "toggle lights",
		},
		{
			"cmd":         "fans",
			"description": "toggle fans",
		},
		{
			"cmd":         "fillWaterTank",
			"description": "fills the water tank",
		},
		{
			"cmd":         "waterPump",
			"description": "toggle water pump",
		},
		{
			"cmd":         "mixValve",
			"description": "toggle mix valve",
		},
		{
			"cmd":         "feedValve",
			"description": "toggle feed valve",
		},
		{
			"cmd":         "stirNutrients",
			"description": "toggle stir nutrients",
		},
		{
			"cmd":         "nutrientOne",
			"description": "toggle nutrient pump 1",
		},
		{
			"cmd":         "nutrientTwo",
			"description": "toggle nutrient pump 2",
		},
		{
			"cmd":         "nutrientThree",
			"description": "toggle nutrient pump 3",
		},
	}
}

// makeRequest is an internal method to change state
func (p *ParticleIO) makeRequest(u, m, action string) error {
	client := &http.Client{}

	data := url.Values{}
	data.Add("args", action)
	data.Add("access_token", p.AccessToken)

	req, err := http.NewRequest(m, u, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return ErrRequestCreationFailed
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		return ErrStateChangeRequestFailed
	}
	if resp.StatusCode != 200 {
		return ErrStateChangeRequestFailed
	}
	return nil
}

// Name Implements driver interface
func (p *ParticleIO) Name() string {
	return "ParticleIO"
}

// Description Implements driver interface
func (p *ParticleIO) Description() string {
	return "ParticleIO gives you everything you need to securely and reliably connect your IoT devices to the web."
}

// New returns a new particle io driver
func New() *ParticleIO {
	p := &ParticleIO{}
	p.UUID = uuid.NewV4()
	return p
}
