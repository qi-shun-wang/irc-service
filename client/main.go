package main

// func ping(a string) {
// 	addr, err := net.ResolveUDPAddr("udp", a)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	c, err := net.DialUDP("udp", nil, addr)
// 	for {
// 		c.Write([]byte("hello, world\n"))
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func main() {
// 	ping(srvAddr)
// }

import (
	"log"

	"github.com/dustin/go-coap"
)

func main() {

	req := coap.Message{
		Type:      coap.NonConfirmable,
		Code:      coap.GET,
		MessageID: 12345,
	}

	req.AddOption(coap.Observe, 1)
	req.SetPathString("/cmd")

	c, err := coap.Dial("udp", "localhost:5683")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	rv, err := c.Send(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	for err == nil {
		if rv != nil {
			if err != nil {
				log.Fatalf("Error receiving: %v", err)
			}
			log.Printf("Got %s", rv.Payload)
		}
		rv, err = c.Receive()

	}
	log.Printf("Done...\n")

}