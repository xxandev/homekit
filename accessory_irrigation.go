package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryIrrigation struct
type AccessoryIrrigation struct {
	*accessory.A
	Valve *service.Valve
}

func (acc *AccessoryIrrigation) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryIrrigation) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryIrrigation) SetID(id uint64) {
	acc.A.Id = id
}
func (acc *AccessoryIrrigation) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryIrrigation) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryIrrigation) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryIrrigation return *Irrigation.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessoryIrrigation(info accessory.Info, args ...interface{}) *AccessoryIrrigation {
	acc := AccessoryIrrigation{}
	acc.A = accessory.New(info, accessory.TypeSprinkler)
	acc.Valve = service.NewValve()
	acc.Valve.ValveType.SetValue(1)
	acc.AddS(acc.Valve.S)
	return &acc
}

//NewAccIrrigation return *Irrigation.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccIrrigation(id uint64, info accessory.Info, args ...interface{}) *AccessoryIrrigation {
	acc := AccessoryIrrigation{}
	acc.A = accessory.New(info, accessory.TypeSprinkler)
	acc.Valve = service.NewValve()
	acc.Valve.ValveType.SetValue(1)
	acc.AddS(acc.Valve.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessoryIrrigation) OnValuesRemoteUpdates(fn func()) {
	acc.Valve.Active.OnValueRemoteUpdate(func(int) { fn() })
	// acc.Valve.InUse.OnValueRemoteUpdate(func(int) { fn() })
}
