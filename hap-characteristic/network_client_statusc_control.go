package hapc

import "github.com/brutella/hap/characteristic"

//TypeNetworkClientStatusControl - 0000020D-0000-1000-8000-0026BB765291
const TypeNetworkClientStatusControl string = "20D"

//NetworkClientStatusControl - Formats TLV8
type NetworkClientStatusControl struct {
	*characteristic.Bytes
}

//NewNetworkClientStatusControl return *NetworkClientStatusControl
func NewNetworkClientStatusControl() *NetworkClientStatusControl {
	char := characteristic.NewBytes(TypeNetworkClientStatusControl)
	char.Format = characteristic.FormatTLV8
	char.Permissions = []string{
		characteristic.PermissionRead,
		characteristic.PermissionWrite,
		characteristic.PermissionEvents,
	}

	char.SetValue([]byte{})

	return &NetworkClientStatusControl{char}
}
