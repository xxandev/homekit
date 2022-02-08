package homekit

import (
	haps "github.com/alpr777/homekit/hap-service"
	"github.com/brutella/hc/accessory"
)

//AccessoryWifiRouter struct
//
//https://github.com/homebridge/HAP-NodeJS/blob/master/src/accessories/Wi-FiRouter_accessory.ts
type AccessoryWifiRouter struct {
	*accessory.Accessory
	Router *haps.WiFiRouter
}

func (acc *AccessoryWifiRouter) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryWifiRouter) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryWifiRouter) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryWifiRouter) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryWifiRouter) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryWifiRouter returns AccessorySwitch (args... are not used)
func NewAccessoryWifiRouter(info accessory.Info, args ...interface{}) *AccessoryWifiRouter {
	acc := AccessoryWifiRouter{}
	acc.Accessory = accessory.New(info, AccessoryTypeWiFiRouter)
	acc.Router = haps.NewWiFiRouter()
	acc.AddService(acc.Router.Service)
	return &acc
}

func (acc *AccessoryWifiRouter) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessoryWifiRouter) OnExample()                      {}
