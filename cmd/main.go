package main

import (
	"fmt"
	"github.com/inercia/water/tuntap"
	"github.com/inercia/water/util"
)

const BUFFERSIZE = 1522

func main() {
	ifce, err := tuntap.NewTAP("")
	if err != nil {
		fmt.Printf("ERROR: when initializing tun/tap device: %s\n", err)
		return
	}

	buffer := make([]byte, BUFFERSIZE)
	for {
		_, err = ifce.Read(buffer)
		if err != nil {
			fmt.Printf("ERROR: when reading from tun/tap device: %s\n", err)
			break
		}

		ethertype := util.MACEthertype(buffer)
		fmt.Printf("Ethertype:      %v\n", ethertype)
		if ethertype == util.IPv4 {
			packet := util.MACPayload(buffer)
			if util.IsIPv4(packet) {
				fmt.Printf("Source:      %v [%v]\n", util.MACSource(buffer), util.IPv4Source(packet))
				fmt.Printf("Destination: %v [%v]\n", util.MACDestination(buffer), util.IPv4Destination(packet))
				fmt.Printf("Protocol:    %v\n\n", util.IPv4Protocol(packet))
			}
		}
	}
}
