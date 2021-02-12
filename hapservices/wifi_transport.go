package hapservices

import (
	"github.com/alpr777/homekit/hapcharacteristics"
	"github.com/brutella/hc/service"
)

//TypeWifiTransport - 00000203-0000-1000-8000-0000022A
const TypeWifiTransport string = "22A"

//WiFiTransport (+CurrentTransport, +WifiCapabilities, WifiConfigurationControl)
type WiFiTransport struct {
	*service.Service

	CurrentTransport         *hapcharacteristics.CurrentTransport
	WifiCapabilities         *hapcharacteristics.WifiCapabilities
	WifiConfigurationControl *hapcharacteristics.WifiConfigurationControl
}

//NewWifiTransport return *WiFiTransport
func NewWifiTransport() *WiFiTransport {
	svc := WiFiTransport{}
	svc.Service = service.New(TypeWiFiSatellite)

	svc.CurrentTransport = hapcharacteristics.NewCurrentTransport()
	svc.AddCharacteristic(svc.CurrentTransport.Characteristic)

	svc.WifiCapabilities = hapcharacteristics.NewWifiCapabilities()
	svc.AddCharacteristic(svc.WifiCapabilities.Characteristic)

	svc.WifiConfigurationControl = hapcharacteristics.NewWifiConfigurationControl()
	svc.AddCharacteristic(svc.WifiConfigurationControl.Characteristic)

	return &svc
}
