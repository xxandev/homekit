package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryFaucet struct
type AccessoryFaucet struct {
	*accessory.A
	// Faucet *service.Faucet
	Valve *service.Valve
}

func (acc *AccessoryFaucet) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryFaucet) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryFaucet) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryFaucet) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryFaucet) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryFaucet) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryFaucet return *Faucet
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessoryFaucet(info accessory.Info, args ...interface{}) *AccessoryFaucet {
	acc := AccessoryFaucet{}
	acc.A = accessory.New(info, accessory.TypeFaucet)
	// acc.Faucet = service.NewFaucet()
	acc.Valve = service.NewValve()

	acc.Valve.ValveType.SetValue(0)

	// acc.AddService(acc.Faucet.Service)
	acc.AddS(acc.Valve.S)
	return &acc
}

//NewAccFaucet return *Faucet
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccFaucet(id uint64, info accessory.Info, args ...interface{}) *AccessoryFaucet {
	acc := AccessoryFaucet{}
	acc.A = accessory.New(info, accessory.TypeFaucet)
	// acc.Faucet = service.NewFaucet()
	acc.Valve = service.NewValve()

	acc.Valve.ValveType.SetValue(0)

	// acc.AddService(acc.Faucet.Service)
	acc.AddS(acc.Valve.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessoryFaucet) OnValuesRemoteUpdates(fn func()) {
	acc.Valve.Active.OnValueRemoteUpdate(func(int) { fn() })
	// acc.Valve.InUse.OnValueRemoteUpdate(func(int) { fn() })
}
