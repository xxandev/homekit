package haps

import (
	"github.com/brutella/hap/service"
	hapc "github.com/xxandev/homekit/hap-characteristic"
)

//TypeWiFiSatellite - 0000020F-0000-1000-8000-0026BB765291
const TypeWiFiSatellite string = "20F"

//WiFiSatellite
//	◈ WiFiSatelliteStatus
type WiFiSatellite struct {
	*service.S
	WiFiSatelliteStatus *hapc.WiFiSatelliteStatus
}

//NewWiFiSatellite return *WiFiSatellite
func NewWiFiSatellite() *WiFiSatellite {
	svc := WiFiSatellite{}
	svc.S = service.New(TypeWiFiSatellite)

	svc.WiFiSatelliteStatus = hapc.NewWiFiSatelliteStatus()
	svc.AddC(svc.WiFiSatelliteStatus.C)

	return &svc
}
