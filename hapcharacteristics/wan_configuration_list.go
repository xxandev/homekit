package hapcharacteristics

import "github.com/brutella/hc/characteristic"

//TypeWANConfigurationList -
const TypeWANConfigurationList string = "211"

//WANConfigurationList -
type WANConfigurationList struct {
	*characteristic.Bytes
}

//NewWANConfigurationList -
func NewWANConfigurationList() *WANConfigurationList {
	char := characteristic.NewBytes(TypeWANConfigurationList)
	char.Format = characteristic.FormatTLV8
	char.Perms = []string{characteristic.PermRead, characteristic.PermEvents}

	char.SetValue([]byte{})

	return &WANConfigurationList{char}
}
