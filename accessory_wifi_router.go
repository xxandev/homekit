package homekit

import (
	"github.com/alpr777/homekit/hapservices"
	"github.com/brutella/hc/accessory"
)

//AccessoryWifiRouter struct
//
//https://github.com/homebridge/HAP-NodeJS/blob/master/src/accessories/Wi-FiRouter_accessory.ts
type AccessoryWifiRouter struct {
	*accessory.Accessory
	Router *hapservices.WiFiRouter
}

//NewAccessoryWifiRouter returns AccessorySwitch (args... are not used)
func NewAccessoryWifiRouter(info accessory.Info, args ...interface{}) *AccessoryWifiRouter {
	acc := AccessoryWifiRouter{}
	acc.Accessory = accessory.New(info, AccessoryTypeWiFiRouter)
	acc.Router = hapservices.NewWiFiRouter()
	acc.AddService(acc.Router.Service)
	return &acc
}
