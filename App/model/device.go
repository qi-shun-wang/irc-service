package model

import (
	"encoding/json"
)

//Device struct that used to describe  current KOD Device plain object.
type Device struct {
	Name    string
	Address string
}

func (device Device) ToJSONString() string {

	bytes, err := json.Marshal(device)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

//Prepare current KOD device Info.
func Prepare() Device {
	device := Device{}
	device.Name = "KOD"
	device.Address = "www.google.com"
	return device
}
