package hapcharacteristics

import "github.com/brutella/hc/characteristic"

//TypeNetworkClientStatusControl -
const TypeNetworkClientStatusControl string = "20D"

//NetworkClientStatusControl -
type NetworkClientStatusControl struct {
	*characteristic.Bytes
}

//NewNetworkClientStatusControl -
func NewNetworkClientStatusControl() *NetworkClientStatusControl {
	char := characteristic.NewBytes(TypeNetworkClientStatusControl)
	char.Format = characteristic.FormatTLV8
	char.Perms = []string{characteristic.PermRead, characteristic.PermWrite, characteristic.PermEvents}

	char.SetValue([]byte{})

	return &NetworkClientStatusControl{char}
}
