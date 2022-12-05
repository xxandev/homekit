package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySwitch struct
type AccessorySwitch struct {
	*accessory.A
	Switch *service.Switch
}

func (acc *AccessorySwitch) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySwitch) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySwitch) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySwitch) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySwitch) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySwitch) GetAccessory() *accessory.A {
	return acc.A
}

// NewAccessorySwitch returns *Switch.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessorySwitch(info accessory.Info, args ...interface{}) *AccessorySwitch {
	acc := AccessorySwitch{}
	acc.A = accessory.New(info, accessory.TypeSwitch)
	acc.Switch = service.NewSwitch()
	acc.AddS(acc.Switch.S)
	return &acc
}

// NewAccSwitch returns *Switch.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccSwitch(id uint64, info accessory.Info, args ...interface{}) *AccessorySwitch {
	acc := AccessorySwitch{}
	acc.A = accessory.New(info, accessory.TypeSwitch)
	acc.Switch = service.NewSwitch()
	acc.AddS(acc.Switch.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessorySwitch) OnValuesRemoteUpdates(fn func()) {
	acc.Switch.On.OnValueRemoteUpdate(func(bool) { fn() })
}
