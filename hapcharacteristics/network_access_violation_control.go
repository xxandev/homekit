package hapcharacteristics

import "github.com/brutella/hc/characteristic"

//TypeNetworkAccessViolationControl -
const TypeNetworkAccessViolationControl string = "21F"

//NetworkAccessViolationControl -
type NetworkAccessViolationControl struct {
	*characteristic.Bytes
}

//NewNetworkAccessViolationControl -
func NewNetworkAccessViolationControl() *NetworkAccessViolationControl {
	char := characteristic.NewBytes(TypeNetworkAccessViolationControl)
	char.Format = characteristic.FormatTLV8
	char.Perms = []string{characteristic.PermRead, characteristic.PermWrite, characteristic.PermEvents}

	char.SetValue([]byte{})

	return &NetworkAccessViolationControl{char}
}
