package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryGate struct
type AccessoryGate struct {
	*accessory.Accessory
	GarageDoorOpener *service.GarageDoorOpener
}

//NewAccessoryGate return AccessoryGate (args... are not used)
func NewAccessoryGate(info accessory.Info, args ...interface{}) *AccessoryGate {
	acc := AccessoryGate{}
	acc.Accessory = accessory.New(info, accessory.TypeGarageDoorOpener)
	acc.GarageDoorOpener = service.NewGarageDoorOpener()
	acc.AddService(acc.GarageDoorOpener.Service)
	return &acc
}
