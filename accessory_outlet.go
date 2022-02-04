package homekit

import (
	"fmt"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryOutlet struct
type AccessoryOutlet struct {
	*accessory.Accessory
	Outlet *service.Outlet
}

func (acc *AccessoryOutlet) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryOutlet) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryOutlet) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryOutlet) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryOutlet) GetAccessory() *accessory.Accessory {
	return acc.Accessory
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

func (acc *AccessoryOutlet) OnValuesRemoteUpdatesPrint() {
	acc.Outlet.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%T - %s] remote update on: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
