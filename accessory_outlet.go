package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryOutlet struct
type AccessoryOutlet struct {
	*accessory.A
	Outlet *service.Outlet
}

func (acc *AccessoryOutlet) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryOutlet) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryOutlet) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryOutlet) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryOutlet) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryOutlet) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryOutlet return *Outlet.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessoryOutlet(info accessory.Info, args ...interface{}) *AccessoryOutlet {
	acc := AccessoryOutlet{}
	acc.A = accessory.New(info, accessory.TypeOutlet)
	acc.Outlet = service.NewOutlet()
	acc.Outlet.OutletInUse.SetValue(true)
	acc.AddS(acc.Outlet.S)
	return &acc
}

//NewAccOutlet return *Outlet.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccOutlet(id uint64, info accessory.Info, args ...interface{}) *AccessoryOutlet {
	acc := AccessoryOutlet{}
	acc.A = accessory.New(info, accessory.TypeOutlet)
	acc.Outlet = service.NewOutlet()
	acc.Outlet.OutletInUse.SetValue(true)
	acc.AddS(acc.Outlet.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessoryOutlet) OnValuesRemoteUpdates(fn func()) {
	acc.Outlet.On.OnValueRemoteUpdate(func(bool) { fn() })
	// acc.Outlet.OutletInUse.OnValueRemoteUpdate(func(bool) { fn() })
}
