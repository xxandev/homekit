package homekit

import (
	haps "github.com/alpr777/homekit/hap-service"
	"github.com/brutella/hc/accessory"
)

//AccessoryFanV2Multifunc struct
type AccessoryFanV2Multifunc struct {
	*accessory.Accessory
	FanV2 *haps.FanV2Multifunc
}

//NewAccessoryFanV2Multifunc return AccessoryFanV2Multifunc (args... are not used)
func NewAccessoryFanV2Multifunc(info accessory.Info, args ...interface{}) *AccessoryFanV2Multifunc {
	acc := AccessoryFanV2Multifunc{}
	acc.Accessory = accessory.New(info, accessory.TypeFan)
	acc.FanV2 = haps.NewFanV2Multifunc()
	acc.AddService(acc.FanV2.Service)
	return &acc
}
