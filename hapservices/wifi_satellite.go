package hapservices

import (
	"github.com/alpr777/homekit/hapcharacteristics"
	"github.com/brutella/hc/service"
)

//TypeWiFiSatellite -
const TypeWiFiSatellite string = "20A"

//WiFiSatellite -
type WiFiSatellite struct {
	*service.Service
	WiFiSatelliteStatus *hapcharacteristics.WiFiSatelliteStatus
}

//NewWiFiSatellite -
func NewWiFiSatellite() *WiFiSatellite {
	svc := WiFiSatellite{}
	svc.Service = service.New(TypeWiFiSatellite)

	svc.WiFiSatelliteStatus = hapcharacteristics.NewWiFiSatelliteStatus()
	svc.AddCharacteristic(svc.WiFiSatelliteStatus.Characteristic)

	return &svc
}
