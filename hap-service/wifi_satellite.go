package haps

import (
	hapc "github.com/alpr777/homekit/hap-characteristic"
	"github.com/brutella/hc/service"
)

//TypeWiFiSatellite - 0000020F-0000-1000-8000-0026BB765291
const TypeWiFiSatellite string = "20F"

//WiFiSatellite
//	◈ WiFiSatelliteStatus
type WiFiSatellite struct {
	*service.Service
	WiFiSatelliteStatus *hapc.WiFiSatelliteStatus
}

//NewWiFiSatellite return *WiFiSatellite
func NewWiFiSatellite() *WiFiSatellite {
	svc := WiFiSatellite{}
	svc.Service = service.New(TypeWiFiSatellite)

	svc.WiFiSatelliteStatus = hapc.NewWiFiSatelliteStatus()
	svc.AddCharacteristic(svc.WiFiSatelliteStatus.Characteristic)

	return &svc
}
