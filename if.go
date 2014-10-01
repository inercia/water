package water

import (
	"os"
)

// Interface is a TUN/TAP interface.
type Interface struct {
	isTAP bool
	file  *os.File
	name  string
}

// Returns true if ifce is a TUN interface, otherwise returns false;
func (ifce *Interface) IsTUN() bool {
	return !ifce.isTAP
}

// Returns true if ifce is a TAP interface, otherwise returns false;
func (ifce *Interface) IsTAP() bool {
	return ifce.isTAP
}

// Returns the interface name of ifce, e.g. tun0, tap1, etc..
func (ifce *Interface) Name() string {
	return ifce.name
}

// Implement io.Writer interface.
func (ifce *Interface) Write(p []byte) (n int, err error) {
	n, err = ifce.file.Write(p)
	return
}

// Implement io.Reader interface.
func (ifce *Interface) Read(p []byte) (n int, err error) {
	n, err = ifce.file.Read(p)
	return
}
