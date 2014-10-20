Description
===========

[![Build Status](https://drone.io/github.com/inercia/water/status.png)](https://drone.io/github.com/inercia/water/latest)

`water` is a native Go library for [TUN/TAP](http://en.wikipedia.org/wiki/TUN/TAP) interfaces.

`water` is designed to be simple and efficient. It

* wraps almost only syscalls and uses only Go standard types;
* exposes standard interfaces; plays well with standard packages like `io`, `bufio`, etc..
* does not handle memory management (allocating/destructing slice). It's up to user to decide how to deal with buffers; whether to use GC.

`water/util` has some useful functions to interpret MAC farme headers and IP packet headers. It also contains some constants such as protocol numbers and ethernet frame types.

Installation
------------

```
go get -u github.com/inercia/water/tuntap
go get -u github.com/inercia/water/util
```

Documentation
-------------

[http://godoc.org/github.com/inercia/water](http://godoc.org/github.com/inercia/water)

[http://godoc.org/github.com/inercia/water/util](http://godoc.org/github.com/inercia/water/waterutil)

Example
-------

```go
package main

import (
	"github.com/inercia/water/tuntap"
	"github.com/inercia/water/util"
	"fmt"
)

const BUFFERSIZE = 1522

func main() {
	ifce, err := tuntap.NewTAP("")
	fmt.Printf("%v, %v\n\n", err, ifce)
	buffer := make([]byte, BUFFERSIZE)
	for {
		_, err = ifce.Read(buffer)
		if err != nil {
			break
		}
		ethertype := util.MACEthertype(buffer)
		if ethertype == util.IPv4 {
			packet := util.MACPayload(buffer)
			if waterutil.IsIPv4(packet) {
				fmt.Printf("Source:      %v [%v]\n", util.MACSource(buffer), util.IPv4Source(packet))
				fmt.Printf("Destination: %v [%v]\n", util.MACDestination(buffer), util.IPv4Destination(packet))
				fmt.Printf("Protocol:    %v\n\n", util.IPv4Protocol(packet))
			}
		}
	}
}
```

This piece of code creates a `TAP` interface, and prints some header information for every IPv4 packet. After pull up the `main.go`, you'll need to bring up the interface and assign IP address. All of these need root permission.

```bash
sudo go run main.go
```

```bash
sudo ip link set dev tap0 up
sudo ip addr add 10.0.0.1/24 dev tap0
```

Now, try sending some ICMP broadcast message:
```bash
ping -b 10.0.0.255
```

You'll see the `main.go` print something like:
```
<nil>, &{true 0xf84003f058 tap0}

Source:      42:35:da:af:2b:00 [10.0.0.1]
Destination: ff:ff:ff:ff:ff:ff [10.0.0.255]
Protocol:    1

Source:      42:35:da:af:2b:00 [10.0.0.1]
Destination: ff:ff:ff:ff:ff:ff [10.0.0.255]
Protocol:    1

Source:      42:35:da:af:2b:00 [10.0.0.1]
Destination: ff:ff:ff:ff:ff:ff [10.0.0.255]
Protocol:    1
```

Changes from original source tree
---------------------------------

* Added support for OS X tap devices.

TODO
----

* IPv6 Support in `util`

LICENSE
-------

[BSD 3-Clause License](http://opensource.org/licenses/BSD-3-Clause)

Alternatives
------------

`tuntap`: [https://code.google.com/p/tuntap/](https://code.google.com/p/tuntap/)

