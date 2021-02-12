package hapcharacteristics

import "github.com/brutella/hc/characteristic"

const (
	//ManagedNetworkEnableDisable -
	ManagedNetworkEnableDisable int = 0
	//ManagedNetworkEnableEnable -
	ManagedNetworkEnableEnable int = 1
	//ManagedNetworkEnableUnknown -
	ManagedNetworkEnableUnknown int = 2
)

//TypeManagedNetworkEnable -
const TypeManagedNetworkEnable string = "215"

//ManagedNetworkEnable -
type ManagedNetworkEnable struct {
	*characteristic.Int
}

//NewManagedNetworkEnable -
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
