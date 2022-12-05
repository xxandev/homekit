package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryGate struct
type AccessoryGate struct {
	*accessory.A
	GarageDoorOpener *service.GarageDoorOpener
}

func (acc *AccessoryGate) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryGate) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryGate) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryGate) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryGate) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryGate) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryGate return *GarageDoorOpener.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessoryGate(info accessory.Info, args ...interface{}) *AccessoryGate {
	acc := AccessoryGate{}
	acc.A = accessory.New(info, accessory.TypeGarageDoorOpener)
	acc.GarageDoorOpener = service.NewGarageDoorOpener()
	acc.AddS(acc.GarageDoorOpener.S)
	return &acc
}

//NewAccGate return *GarageDoorOpener.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccGate(id uint64, info accessory.Info, args ...interface{}) *AccessoryGate {
	acc := AccessoryGate{}
	acc.A = accessory.New(info, accessory.TypeGarageDoorOpener)
	acc.GarageDoorOpener = service.NewGarageDoorOpener()
	acc.AddS(acc.GarageDoorOpener.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessoryGate) OnValuesRemoteUpdates(fn func()) {
	acc.GarageDoorOpener.TargetDoorState.OnValueRemoteUpdate(func(int) { fn() })
	// acc.GarageDoorOpener.ObstructionDetected.OnValueRemoteUpdate(func(bool) { fn() })
}
