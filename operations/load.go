package operations

import (
    "encoding/json"
    "errors"
    "io"
    "os"

    "godel/types"
)

func Load(sources []string) (map[string]types.Device, error) {
    var devices = make(map[string]types.Device)

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
            var device types.Device
            if err = jsonDec.Decode(&device); err == io.EOF {
                break
            } else if err != nil {
                return nil, errors.New("Could not read JSON object in source \"" + source + "\": " + err.Error())
            }
            devices[device.Address] = device
        }
    }

    return devices, nil
}
