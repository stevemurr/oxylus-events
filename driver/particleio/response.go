package particleio

import (
	"fmt"
	"time"
)

/*
POST response
{
	"id":"290018000347353137323334",
	"name":"test_photon",
	"last_app":"",
	"connected":true,
	"return_value":1
}
*/

/*
GET response
{
  "cmd": "VarReturn",
  "name": "lightValue",
  "result": 17,
  "coreInfo": {
    "last_app": "",
    "last_heard": "2017-06-19T00:23:41.545Z",
    "connected": true,
    "last_handshake_at": "2017-06-18T23:06:14.571Z",
    "deviceID": "290018000347353137323334",
    "product_id": 4439
  }
}
*/

// Response data from the particleIO service
// Result contains the sensor value
type Response struct {
	CMD      string   `json:"cmd"`
	Name     string   `json:"name"`
	Result   int      `json:"result"`
	CoreInfo coreInfo `json:"coreInfo"`
}

type coreInfo struct {
	LastApp         string    `json:"last_app"`
	LastHeard       string    `json:"last_heard"`
	Connected       bool      `json:"connected"`
	LastHandshakeAt time.Time `json:"last_handshake_at"`
	DeviceID        string    `json:"deviceID"`
	ProductID       int       `json:"product_id"`
}

func (r *Response) String() string {
	return fmt.Sprintf("%s %s %d", r.CMD, r.Name, r.Result)
}

// OxylusResponse represents our response object that driver responses are converted to
// before they are inserted to our db
type OxylusResponse struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (o *OxylusResponse) String() string {
	return fmt.Sprintf("%s %s", o.Name, o.Value)
}
