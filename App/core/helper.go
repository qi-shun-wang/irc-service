package core

import (
	"log"
	"net"
)

//GetSettingsSystemInfo .
func GetSettingsSystemInfo(ch chan string) {
	c := Commander{}
	result, err := c.OnCmds("settings --user current list system;")

	if err != nil {
		log.Println("GetSettingsSystemInfo ERR:", err)
		// return err.Error()
		ch <- err.Error()
	}

	// log.Print("SUCCESS --> GetSettingsSystemInfo")
	ch <- result
	// return result
}

//GetDeviceName .
func GetDeviceName(ch chan string) {
	cmd := Commander{}
	result, err := cmd.OnCmds("settings --user current get global device_name;")

	if err != nil {
		log.Println("GetDeviceName:", err)
	}
	// log.Print("GetDeviceName  -->:" + result)
	ch <- result
}

//GetBackupIP .
func GetBackupIP(ch chan string) {
	cmd := Commander{}
	result, err := cmd.OnCmds("ip addr show wlan0  | grep 'inet ' | cut -d' ' -f6|cut -d/ -f1;")

	if err != nil {
		log.Println("GetBackupIP:", err)
	}
	// log.Print("GetDeviceName  -->:" + result)
	ch <- result
}

//GetOutboundIP .
func GetOutboundIP(ch chan string) {

	ifaces, err := net.Interfaces()

	// handle err
	if err != nil {
		log.Println("GetOutboundIP:", err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		// handle err
		if err != nil {
			log.Println("GetOutboundIP handler:", err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:

				ip = v.IP
				if !ip.IsLoopback() &&
					ip.To4() != nil &&
					!ip.IsMulticast() &&
					!ip.IsLinkLocalMulticast() &&
					!ip.IsInterfaceLocalMulticast() &&
					// !ip.IsGlobalUnicast() { //&& //two of the ip is gone
					!ip.IsLinkLocalUnicast() { //&&
					//!ip.IsUnspecified() {

					// log.Println("Resolved Host IPNet: " + ip.String())
					ch <- ip.String()
					return
				}

			case *net.IPAddr:
				ip = v.IP
				// log.Println("Resolved Host IPAddr: " + ip.String())
				ch <- ip.String()
			}
		}
	}
}
