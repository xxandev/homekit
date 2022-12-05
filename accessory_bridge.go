package homekit

import (
	"github.com/brutella/hap/accessory"
)

// NewBridge returns a bridge which implements model.Bridge.
type AccessoryBridge struct {
	*accessory.A
}

func (acc *AccessoryBridge) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryBridge) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryBridge) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryBridge) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryBridge) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryBridge) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryBridge returns *Bridge.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessoryBridge(info accessory.Info, args ...interface{}) *AccessoryBridge {
	a := AccessoryBridge{}
	a.A = accessory.New(info, accessory.TypeBridge)
	return &a
}

//NewAccBridge returns *Bridge.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccBridge(id uint64, info accessory.Info, args ...interface{}) *AccessoryBridge {
	a := AccessoryBridge{}
	a.A = accessory.New(info, accessory.TypeBridge)
	a.A.Id = id
	return &a
}

func (acc *AccessoryBridge) OnValuesRemoteUpdates(fn func()) {}
