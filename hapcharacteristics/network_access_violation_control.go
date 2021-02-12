package hapcharacteristics

import "github.com/brutella/hc/characteristic"

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
	char.Perms = []string{characteristic.PermRead, characteristic.PermWrite, characteristic.PermEvents}

	char.SetValue([]byte{})

	return &NetworkAccessViolationControl{char}
}
