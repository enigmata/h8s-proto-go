package controller

import (
	"github.com/enigmata/h8s-proto-go/controller/isy994"
	"github.com/enigmata/h8s-proto-go/device"
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
