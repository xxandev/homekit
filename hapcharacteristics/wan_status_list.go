package hapcharacteristics

import "github.com/brutella/hc/characteristic"

//TypeWANStatusList - 00000212-0000-1000-8000-0026BB765291
const TypeWANStatusList string = "212"

//WANStatusList - Formats TLV8
type WANStatusList struct {
	*characteristic.Bytes
}

//NewWANStatusList return *WANStatusList
func NewWANStatusList() *WANStatusList {
	char := characteristic.NewBytes(TypeWANStatusList)
	char.Format = characteristic.FormatTLV8
	char.Perms = []string{characteristic.PermRead, characteristic.PermEvents}

	char.SetValue([]byte{})

	return &WANStatusList{char}
}
