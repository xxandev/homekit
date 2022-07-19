package homekit

import (
	"github.com/brutella/hap/accessory"
	haps "github.com/xxandev/homekit/hap-service"
)

//AccessoryWifiRouter struct
//
//https://github.com/homebridge/HAP-NodeJS/blob/master/src/accessories/Wi-FiRouter_accessory.ts
type AccessoryWifiRouter struct {
	*accessory.A
	Router *haps.WiFiRouter
}

func (acc *AccessoryWifiRouter) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryWifiRouter) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryWifiRouter) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryWifiRouter) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryWifiRouter) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryWifiRouter) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryWifiRouter returns AccessorySwitch (args... are not used)
func NewAccessoryWifiRouter(info accessory.Info, args ...interface{}) *AccessoryWifiRouter {
	acc := AccessoryWifiRouter{}
	acc.A = accessory.New(info, AccessoryTypeWiFiRouter)
	acc.Router = haps.NewWiFiRouter()
	acc.AddS(acc.Router.S)
	return &acc
}

func (acc *AccessoryWifiRouter) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessoryWifiRouter) OnExample()                      {}
