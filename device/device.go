package device

import (
	"fmt"
)

type Device struct {
	Name         string `json:"name,omitempty"`
	Address      string `json:"address,omitempty"`
	Type         string `json:"type,omitempty"`
	Enabled      bool   `json:"enabled",omitempty"`
	Room         string `json:"room,omitempty"`
	RoomLocation string `json:"roomLocation,omitempty"`
}

func (d Device) Print(indent string) {
	fmt.Printf("%sName:     %s\n", indent, d.Name)
	fmt.Printf("%sAddress:  %s\n", indent, d.Address)
	fmt.Printf("%sType:     %s\n", indent, d.Type)
	fmt.Printf("%sEnabled:  %t\n", indent, d.Enabled)
	fmt.Printf("%sRoom:     %s\n", indent, d.Room)
	fmt.Printf("%sLocation: %s\n", indent, d.RoomLocation)
}
