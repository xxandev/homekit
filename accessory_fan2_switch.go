package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryFanSwitch struct
type AccessoryFan2Switch struct {
	*accessory.Accessory
	Fan2 *service.FanV2
}

//NewAccessoryFanSwitch return AccessoryFanSwitch (args... are not used)
func NewAccessoryFan2Switch(info accessory.Info, args ...interface{}) *AccessoryFan2Switch {
	acc := AccessoryFan2Switch{}
	acc.Accessory = accessory.New(info, accessory.TypeFan)
	acc.Fan2 = service.NewFanV2()
	acc.AddService(acc.Fan2.Service)
	return &acc
}
