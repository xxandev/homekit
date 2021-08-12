package hapc

import "github.com/brutella/hc/characteristic"

//TypeWifiConfigurationControl - 0000021E-0000-1000-8000-0000022D
const TypeWifiConfigurationControl string = "22D"

//WifiConfigurationControl - Formats TLV8
type WifiConfigurationControl struct {
	*characteristic.Bytes
}

//NewWifiConfigurationControl return *WifiConfigurationControl
func NewWifiConfigurationControl() *WifiConfigurationControl {
	char := characteristic.NewBytes(TypeWifiCapabilities)
	char.Format = characteristic.FormatTLV8
	char.Perms = []string{characteristic.PermRead, characteristic.PermWrite, characteristic.PermEvents}

	return &WifiConfigurationControl{char}
}
