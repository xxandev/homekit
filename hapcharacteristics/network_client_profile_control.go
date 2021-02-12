package hapcharacteristics

import "github.com/brutella/hc/characteristic"

//TypeNetworkClientProfileControl -
const TypeNetworkClientProfileControl string = "20C"

//NetworkClientProfileControl -
type NetworkClientProfileControl struct {
	*characteristic.Bytes
}

//NewNetworkClientProfileControl -
func NewNetworkClientProfileControl() *NetworkClientProfileControl {
	char := characteristic.NewBytes(TypeNetworkClientProfileControl)
	char.Format = characteristic.FormatTLV8
	char.Perms = []string{characteristic.PermRead, characteristic.PermWrite, characteristic.PermEvents}

	char.SetValue([]byte{})

	return &NetworkClientProfileControl{char}
}
