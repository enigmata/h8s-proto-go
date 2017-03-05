package operations

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"godel/device"
)

func Load(sources []string) (map[string]device.Device, error) {
	var devices = make(map[string]device.Device)

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
			var dev device.Device
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
