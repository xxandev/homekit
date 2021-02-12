package hapcharacteristics

import "github.com/brutella/hc/characteristic"

const (
	//WiFiSatelliteStatusUnknown - unknown(0)
	WiFiSatelliteStatusUnknown int = 0

	//WiFiSatelliteStatusConnect - connect(1)
	WiFiSatelliteStatusConnect int = 1

	//WiFiSatelliteStatusNotConnect - not connect(2)
	WiFiSatelliteStatusNotConnect int = 2
)

//TypeWiFiSatelliteStatus - 0000021E-0000-1000-8000-0026BB765291
const TypeWiFiSatelliteStatus string = "21E"

//WiFiSatelliteStatus - Formats UINT8
type WiFiSatelliteStatus struct {
	*characteristic.Int
}

//NewWiFiSatelliteStatus return *WiFiSatelliteStatus
func NewWiFiSatelliteStatus() *WiFiSatelliteStatus {
	char := characteristic.NewInt(TypeWiFiSatelliteStatus)
	char.Format = characteristic.FormatUInt8
	char.Perms = []string{characteristic.PermRead, characteristic.PermWrite, characteristic.PermEvents}

	char.SetMinValue(0)
	char.SetMaxValue(2)
	char.SetStepValue(1)
	char.SetValue(0)

	return &WiFiSatelliteStatus{char}
}
