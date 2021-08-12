package hapc

import "github.com/brutella/hc/characteristic"

//TypeNetworkClientProfileControl - 0000020C-0000-1000-8000-0026BB765291
const TypeNetworkClientProfileControl string = "20C"

//NetworkClientProfileControl - Formats TLV8
type NetworkClientProfileControl struct {
	*characteristic.Bytes
}

//NewNetworkClientProfileControl return *NetworkClientProfileControl
func NewNetworkClientProfileControl() *NetworkClientProfileControl {
	char := characteristic.NewBytes(TypeNetworkClientProfileControl)
	char.Format = characteristic.FormatTLV8
	char.Perms = []string{characteristic.PermRead, characteristic.PermWrite, characteristic.PermEvents}

	char.SetValue([]byte{})

	return &NetworkClientProfileControl{char}
}
