package haps

import (
	hapc "github.com/alpr777/homekit/hap-characteristic"
	"github.com/brutella/hc/service"
)

//TypeWifiTransport - 00000203-0000-1000-8000-0000022A
const TypeWiFiTransport string = "22A"

//WiFiTransport (+CurrentTransport, +WifiCapabilities, WifiConfigurationControl)
type WiFiTransport struct {
	*service.Service

	CurrentTransport         *hapc.CurrentTransport
	WifiCapabilities         *hapc.WifiCapabilities
	WifiConfigurationControl *hapc.WifiConfigurationControl
}

//NewWifiTransport return *WiFiTransport
func NewWifiTransport() *WiFiTransport {
	svc := WiFiTransport{}
	svc.Service = service.New(TypeWiFiTransport)

	svc.CurrentTransport = hapc.NewCurrentTransport()
	svc.AddCharacteristic(svc.CurrentTransport.Characteristic)

	svc.WifiCapabilities = hapc.NewWifiCapabilities()
	svc.AddCharacteristic(svc.WifiCapabilities.Characteristic)

	svc.WifiConfigurationControl = hapc.NewWifiConfigurationControl()
	svc.AddCharacteristic(svc.WifiConfigurationControl.Characteristic)

	return &svc
}
