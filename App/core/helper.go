package core

import (
	"log"
	"net"
	"strings"
)

//GetDeviceName .
func GetDeviceName() string {
	c := Commander{}
	result, err := c.OnCmds("settings --user current list global |grep -w device_name")

	if err != nil {
		log.Println("GetDeviceName:", err)
	}
	deviceName := strings.TrimPrefix(result, "device_name=")
	log.Print(deviceName)
	return deviceName
}

//GetOutboundIP .
func GetOutboundIP() string {

	netInterfaceAddresses, err := net.InterfaceAddrs()

	if err != nil {
		log.Println("GetOutboundIP:", err)
	}

	for _, netInterfaceAddress := range netInterfaceAddresses {

		network, ok := netInterfaceAddress.(*net.IPNet)

		if ok && !network.IP.IsLoopback() && network.IP.To4() != nil {

			ip := network.IP.String()

			log.Println("Resolved Host IP: " + ip)

			return ip
		}
	}
	return ""

}
