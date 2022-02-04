package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryOutlet struct
type AccessoryOutlet struct {
	*accessory.Accessory
	Outlet *service.Outlet
}

//NewAccessoryOutlet return AccessoryOutlet (args... are not used)
func NewAccessoryOutlet(info accessory.Info, args ...interface{}) *AccessoryOutlet {
	acc := AccessoryOutlet{}
	acc.Accessory = accessory.New(info, accessory.TypeOutlet)
	acc.Outlet = service.NewOutlet()
	acc.Outlet.OutletInUse.SetValue(true)
	acc.AddService(acc.Outlet.Service)
	return &acc
}

func (acc *AccessoryOutlet) OnValuesRemoteUpdates(fn func()) {
	acc.Outlet.On.OnValueRemoteUpdate(func(bool) { fn() })
	// acc.Outlet.OutletInUse.OnValueRemoteUpdate(func(bool) { fn() })
}
