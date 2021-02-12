package hapcharacteristics

import "github.com/brutella/hc/characteristic"

//TypeWANStatusList -
const TypeWANStatusList string = "212"

//WANStatusList -
type WANStatusList struct {
	*characteristic.Bytes
}

//NewWANStatusList -
func NewWANStatusList() *WANStatusList {
	char := characteristic.NewBytes(TypeWANStatusList)
	char.Format = characteristic.FormatTLV8
	char.Perms = []string{characteristic.PermRead, characteristic.PermEvents}

	char.SetValue([]byte{})

	return &WANStatusList{char}
}
