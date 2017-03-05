package device

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
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

func LoadFromFile(sources []string) (map[string]Device, error) {
	var devices = make(map[string]Device)

	if len(sources) < 1 {
		return devices, errors.New("No sources to load")
	}

	for _, source := range sources {
		file, err := os.Open(source)
		if err != nil {
			return nil, errors.New("Could not open source \"" + source + "\"")
		}
		defer file.Close()

		jsonDec := json.NewDecoder(file)

		for {
			var dev Device
			if err = jsonDec.Decode(&dev); err == io.EOF {
				break
			} else if err != nil {
				return nil, errors.New("Could not read JSON object in source \"" + source + "\": " + err.Error())
			}
			devices[dev.Address] = dev
		}
	}

	return devices, nil
}
