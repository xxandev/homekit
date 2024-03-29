package hapc

import "github.com/brutella/hap/characteristic"

//TypeWANConfigurationList - 00000211-0000-1000-8000-0026BB765291
const TypeWANConfigurationList string = "211"

//WANConfigurationList - Formats TLV8
type WANConfigurationList struct {
	*characteristic.Bytes
}

//NewWANConfigurationList return *WANConfigurationList
func NewWANConfigurationList() *WANConfigurationList {
	char := characteristic.NewBytes(TypeWANConfigurationList)
	char.Format = characteristic.FormatTLV8
	char.Permissions = []string{
		characteristic.PermissionRead,
		characteristic.PermissionEvents,
	}

	char.SetValue([]byte{})

	return &WANConfigurationList{char}
}
