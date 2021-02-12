package hapcharacteristics

import "github.com/brutella/hc/characteristic"

//TypeWifiCapabilities - 0000021E-0000-1000-8000-0000022D
const TypeWifiCapabilities = "22D"

//WifiCapabilities - Formats UINT32
type WifiCapabilities struct {
	*characteristic.Int
}

//NewWifiCapabilities return *WifiCapabilities
func NewWifiCapabilities() *WifiCapabilities {
	char := characteristic.NewInt(TypeWifiCapabilities)
	char.Format = characteristic.FormatUInt32
	char.Perms = []string{characteristic.PermRead}

	char.SetValue(1)
	return &WifiCapabilities{char}
}
