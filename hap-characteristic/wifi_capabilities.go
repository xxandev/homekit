package hapc

import "github.com/brutella/hap/characteristic"

//TypeWifiCapabilities - 0000021E-0000-1000-8000-0000022D
//const TypeWifiCapabilities = "22D"
const TypeWifiCapabilities = "22C"

//WifiCapabilities - Formats UINT32
type WifiCapabilities struct {
	*characteristic.Int
}

//NewWifiCapabilities return *WifiCapabilities
func NewWifiCapabilities() *WifiCapabilities {
	char := characteristic.NewInt(TypeWifiCapabilities)
	char.Format = characteristic.FormatUInt32
	char.Permissions = []string{characteristic.PermissionRead}

	char.SetValue(1)
	return &WifiCapabilities{char}
}
