package haps

import (
	"github.com/brutella/hap/service"
	hapc "github.com/xxandev/homekit/hap-characteristic"
)

//TypeWifiTransport - 00000203-0000-1000-8000-0000022A
const TypeWiFiTransport string = "22A"

//WiFiTransport
//	◈ CurrentTransport
//	◈ WifiCapabilities
//	◇ WifiConfigurationControl
type WiFiTransport struct {
	*service.S

	CurrentTransport         *hapc.CurrentTransport
	WifiCapabilities         *hapc.WifiCapabilities
	WifiConfigurationControl *hapc.WifiConfigurationControl
}

//NewWifiTransport return *WiFiTransport
func NewWifiTransport() *WiFiTransport {
	svc := WiFiTransport{}
	svc.S = service.New(TypeWiFiTransport)

	svc.CurrentTransport = hapc.NewCurrentTransport()
	svc.AddC(svc.CurrentTransport.C)

	svc.WifiCapabilities = hapc.NewWifiCapabilities()
	svc.AddC(svc.WifiCapabilities.C)

	svc.WifiConfigurationControl = hapc.NewWifiConfigurationControl()
	svc.AddC(svc.WifiConfigurationControl.C)

	return &svc
}
