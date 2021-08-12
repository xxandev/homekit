package hapc

import "github.com/brutella/hc/characteristic"

//TypeSupportedRouterConfiguration - 00000210-0000-1000-8000-0026BB765291
const TypeSupportedRouterConfiguration string = "210"

//SupportedRouterConfiguration - Formats TLV8
type SupportedRouterConfiguration struct {
	*characteristic.Bytes
}

//NewSupportedRouterConfiguration return *SupportedRouterConfiguration
func NewSupportedRouterConfiguration() *SupportedRouterConfiguration {
	char := characteristic.NewBytes(TypeSupportedRouterConfiguration)
	char.Format = characteristic.FormatTLV8
	char.Perms = []string{characteristic.PermRead}

	char.SetValue([]byte{})

	return &SupportedRouterConfiguration{char}
}
