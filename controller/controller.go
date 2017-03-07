package controller

import (
	"godel/controller/isy994"
	"godel/device"
)

type ControllerClient interface {
	Devices() map[string]device.Device
}

const (
	ISY994 string = "isy994"
)

func NewClient(name string) ControllerClient {
	switch name {
	case ISY994:
		return isy994.NewClient()
	default:
		return nil
	}
	return nil
}
