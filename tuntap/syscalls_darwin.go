// +build darwin

package tuntap

/*
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"github.com/inercia/kernctl"
	"os"
	"syscall"
	"unsafe"
)

const UTUN_CONTROL_NAME = "com.apple.net.utun_control"
const UTUN_OPT_IFNAME = 2

var ERROR_NOT_DEVICE_FOUND = errors.New("could not obtain valid tun/tap device")

// Create a new TAP interface whose name is ifName.
// If ifName is empty, a default name (tap0, tap1, ... ) will be assigned.
// ifName should not exceed 16 bytes.
func NewTAP(ifName string) (ifce *TunTap, err error) {
	name, file, err := createInterface()
	if err != nil {
		return nil, err
	}
	ifce = &TunTap{isTAP: true, file: file, name: name}
	return
}

// Create a new TUN interface whose name is ifName.
// If ifName is empty, a default name (tap0, tap1, ... ) will be assigned.
// ifName should not exceed 16 bytes.
func NewTUN(ifName string) (ifce *TunTap, err error) {
	name, file, err := createInterface()
	if err != nil {
		return nil, err
	}
	ifce = &TunTap{isTAP: false, file: file, name: name}
	return
}

func createInterface() (createdIFName string, file *os.File, err error) {
	file = nil
	err = ERROR_NOT_DEVICE_FOUND

	var readBufLen C.int = 20
	var readBuf = C.CString("                    ")
	defer C.free(unsafe.Pointer(readBuf))

	for utunnum := 0; utunnum < 255; utunnum++ {
		conn := kernctl.NewConnByName(UTUN_CONTROL_NAME)
		conn.UnitId = uint32(utunnum + 1)
		conn.Connect()

		_, _, gserr := syscall.Syscall6(syscall.SYS_GETSOCKOPT,
			uintptr(conn.Fd),
			uintptr(kernctl.SYSPROTO_CONTROL), uintptr(UTUN_OPT_IFNAME),
			uintptr(unsafe.Pointer(readBuf)), uintptr(unsafe.Pointer(&readBufLen)), 0)
		if gserr == 0 {
			createdIFName = C.GoStringN(readBuf, C.int(readBufLen))
			file = os.NewFile(uintptr(conn.Fd), createdIFName)
			err = nil
			break
		}
	}

	return createdIFName, file, err
}
