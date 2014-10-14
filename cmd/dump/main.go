package main

import (
	"fmt"
	"github.com/inercia/water/tuntap"
	"github.com/inercia/water/util"
)

const BUFFERSIZE = 1522

func main() {
	// you can run this simple test (in OS X):
	//  $ sudo run cmd/main.go
	// (get the device name, <DEV>, from the output)
	//  $ sudo ifconfig <DEV> 10.9.0.1 10.9.255.255
	//  $ sudo ifconfig <DEV> up
	//  $ ping -b <DEV> 10.9.0.8

	ifce, err := tuntap.NewTAP("")
	if err != nil {
		fmt.Printf("ERROR: when initializing tun/tap device: %s\n", err)
		return
	}
	fmt.Printf("Device: %s\n", ifce.Name())

	buffer := make([]byte, BUFFERSIZE)
	for {
		_, err = ifce.Read(buffer)
		if err != nil {
			fmt.Printf("ERROR: when reading from tun/tap device: %s\n", err)
			break
		}

		ethertype := util.MACEthertype(buffer)
		packet := util.MACPayload(buffer)
		fmt.Printf("Ethertype:      %v\n", ethertype)
		fmt.Printf("Source:      %v [%v]\n", util.MACSource(buffer), util.IPv4Source(packet))
		fmt.Printf("Destination: %v [%v]\n", util.MACDestination(buffer), util.IPv4Destination(packet))

		if ethertype == util.IPv4 {
			if util.IsIPv4(packet) {
				fmt.Printf("Protocol:    %v\n\n", util.IPv4Protocol(packet))
			}
		}
	}
}
