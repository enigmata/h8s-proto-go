package types

import (
    "fmt"
)

type Device struct {
    Name string          `json:"name,omitempty"`
    Address string       `json:"address,omitempty"`
    Type string          `json:"type,omitempty"`
    Enabled bool         `json:"enabled",omitempty"`
    Room string          `json:"room,omitempty"`
    RoomLocation string  `json:"roomLocation,omitempty"`
}


func PrintDevice(dev Device, indent string) {
    fmt.Printf("%sName:     %s\n", indent, dev.Name)
    fmt.Printf("%sAddress:  %s\n", indent, dev.Address)
    fmt.Printf("%sType:     %s\n", indent, dev.Type)
    fmt.Printf("%sEnabled:  %t\n", indent, dev.Enabled)
    fmt.Printf("%sRoom:     %s\n", indent, dev.Room)
    fmt.Printf("%sLocation: %s\n", indent, dev.RoomLocation)
}
