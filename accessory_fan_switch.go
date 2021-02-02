package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryFanSwitch struct
type AccessoryFanSwitch struct {
	*accessory.Accessory
	Fan *service.Fan
}

//NewAccessoryFanSwitch return AccessoryFanSwitch (args... are not used)
func NewAccessoryFanSwitch(info accessory.Info, args ...interface{}) *AccessoryFanSwitch {
	acc := AccessoryFanSwitch{}
	acc.Accessory = accessory.New(info, accessory.TypeFan)
	acc.Fan = service.NewFan()
	acc.AddService(acc.Fan.Service)
	return &acc
}
