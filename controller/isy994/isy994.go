package isy994

import (
	"github.com/enigmata/h8s-proto-go/device"
)

type ISY994ControllerClient struct {
	devices map[string]device.Device
}

func (c *ISY994ControllerClient) Devices() map[string]device.Device {
	var devices = make(map[string]device.Device)

	if c.devices == nil {
		c.devices = make(map[string]device.Device)

		c.devices["dummy"] = device.Device{
			Name:         "dummy",
			Address:      "address",
			Type:         "type",
			Enabled:      true,
			Room:         "room",
			RoomLocation: "location in room",
		}
	}

	for k, v := range c.devices {
		devices[k] = v
	}

	return devices
}

func NewClient() *ISY994ControllerClient {
	return &ISY994ControllerClient{}
}
