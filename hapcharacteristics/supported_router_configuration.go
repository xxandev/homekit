package hapcharacteristics

import "github.com/brutella/hc/characteristic"

//TypeSupportedRouterConfiguration -
const TypeSupportedRouterConfiguration string = "210"

//SupportedRouterConfiguration -
type SupportedRouterConfiguration struct {
	*characteristic.Bytes
}

//NewSupportedRouterConfiguration -
func NewSupportedRouterConfiguration() *SupportedRouterConfiguration {
	char := characteristic.NewBytes(TypeSupportedRouterConfiguration)
	char.Format = characteristic.FormatTLV8
	char.Perms = []string{characteristic.PermRead, characteristic.PermWrite, characteristic.PermEvents}

	char.SetValue([]byte{})

	return &SupportedRouterConfiguration{char}
}
