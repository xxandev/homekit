package hapc

import "github.com/brutella/hap/characteristic"

//TypeNetworkAccessViolationControl - 0000021F-0000-1000-8000-0026BB765291
const TypeNetworkAccessViolationControl string = "21F"

//NetworkAccessViolationControl - Formats TLV8
type NetworkAccessViolationControl struct {
	*characteristic.Bytes
}

//NewNetworkAccessViolationControl return *NetworkAccessViolationControl
func NewNetworkAccessViolationControl() *NetworkAccessViolationControl {
	char := characteristic.NewBytes(TypeNetworkAccessViolationControl)
	char.Format = characteristic.FormatTLV8
	char.Permissions = []string{
		characteristic.PermissionRead,
		characteristic.PermissionWrite,
		characteristic.PermissionEvents,
	}

	char.SetValue([]byte{})

	return &NetworkAccessViolationControl{char}
}
