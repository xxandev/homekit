package homekit

import (
	"github.com/alpr777/homekit/hapservices"
	"github.com/brutella/hc/accessory"
)

//AccessoryWifiSatellite struct
//
//https://github.com/homebridge/HAP-NodeJS/blob/106d0ead82eb02b81d3f58f1fbbbc975c52c8857/src/accessories/Wi-FiSatellite_accessory.ts
type AccessoryWifiSatellite struct {
	*accessory.Accessory
	Satellite *hapservices.WiFiSatellite
}

//NewAccessoryWifiSatellite returns AccessorySwitch (args... are not used)
func NewAccessoryWifiSatellite(info accessory.Info, args ...interface{}) *AccessoryWifiSatellite {
	acc := AccessoryWifiSatellite{}
	acc.Accessory = accessory.New(info, AccessoryTypeWiFiRouter)
	acc.Satellite = hapservices.NewWiFiSatellite()

	acc.AddService(acc.Satellite.Service)
	return &acc
}
