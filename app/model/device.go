package model

import (
	helper "IRCService/app/core"
	"encoding/json"
)

//Device struct that used to describe  current KOD Device plain object.
type Device struct {
	Name          *string
	Address       *string
	BackupAddress *string
	Settings      *string
}

//ToJSONString method that convert Device model into JSON string .
func (device Device) ToJSONString() string {

	bytes, err := json.Marshal(device)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

//Prepare current KOD device Info.
func Prepare() Device {
	cName := make(chan string)
	cIP := make(chan string)
	cBackupIP := make(chan string)
	cSetting := make(chan string)
	device := Device{}
	go func() {
		helper.GetDeviceName(cName)
		helper.GetBackupIP(cBackupIP)
		helper.GetOutboundIP(cIP)
		helper.GetSettingsSystemInfo(cSetting)
	}()
	for device.Name == nil || device.Settings == nil || device.Address == nil || device.BackupAddress == nil {
		select {
		case name := <-cName:
			device.Name = &name
		case address := <-cIP:
			device.Address = &address
		case backupAddress := <-cBackupIP:
			device.BackupAddress = &backupAddress
		case setting := <-cSetting:
			device.Settings = &setting
		}
	}
	return device
}
