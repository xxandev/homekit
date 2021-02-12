package hapcharacteristics

import "github.com/brutella/hc/characteristic"

const (
	//WiFiSatelliteStatusUnknown -
	WiFiSatelliteStatusUnknown int = 0

	//WiFiSatelliteStatusConnect -
	WiFiSatelliteStatusConnect int = 1

	//WiFiSatelliteStatusNotConnect -
	WiFiSatelliteStatusNotConnect int = 2
)

//TypeWiFiSatelliteStatus -
const TypeWiFiSatelliteStatus string = "21E"

//WiFiSatelliteStatus -
type WiFiSatelliteStatus struct {
	*characteristic.Int
}

//NewWiFiSatelliteStatus -
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
