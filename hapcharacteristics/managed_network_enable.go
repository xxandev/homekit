package hapcharacteristics

import "github.com/brutella/hc/characteristic"

const (
	//ManagedNetworkEnableDisable - disable(0)
	ManagedNetworkEnableDisable int = 0

	//ManagedNetworkEnableEnable - enable(1)
	ManagedNetworkEnableEnable int = 1

	//ManagedNetworkEnableUnknown - unknown(2)
	ManagedNetworkEnableUnknown int = 2
)

//TypeManagedNetworkEnable - 00000215-0000-1000-8000-0026BB765291
const TypeManagedNetworkEnable string = "215"

//ManagedNetworkEnable - Formats UINT8
type ManagedNetworkEnable struct {
	*characteristic.Int
}

//NewManagedNetworkEnable return *ManagedNetworkEnable
func NewManagedNetworkEnable() *ManagedNetworkEnable {
	char := characteristic.NewInt(TypeManagedNetworkEnable)
	char.Format = characteristic.FormatUInt8
	char.Perms = []string{characteristic.PermRead, characteristic.PermWrite, characteristic.PermEvents}

	char.SetMinValue(0)
	char.SetMaxValue(1)
	char.SetStepValue(1)
	char.SetValue(0)

	return &ManagedNetworkEnable{char}
}
