package homekit

import (
	"github.com/brutella/hap/accessory"
	haps "github.com/xxandev/homekit/hap-service"
)

//AccessoryFanV2Multifunc struct
type AccessoryFan2RS struct {
	*accessory.A
	Fan2 *haps.Fan2RS
}

func (acc *AccessoryFan2RS) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryFan2RS) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryFan2RS) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryFan2RS) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryFan2RS) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryFan2RS) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryFan2RS return *FanV2RS.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessoryFan2RS(info accessory.Info, args ...interface{}) *AccessoryFan2RS {
	acc := AccessoryFan2RS{}
	acc.A = accessory.New(info, accessory.TypeFan)
	acc.Fan2 = haps.NewFanV2RS()
	acc.AddS(acc.Fan2.S)
	return &acc
}

//NewAccFan2RS return *FanV2RS.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccFan2RS(id uint64, info accessory.Info, args ...interface{}) *AccessoryFan2RS {
	acc := AccessoryFan2RS{}
	acc.A = accessory.New(info, accessory.TypeFan)
	acc.Fan2 = haps.NewFanV2RS()
	acc.AddS(acc.Fan2.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessoryFan2RS) OnValuesRemoteUpdates(fn func()) {
	acc.Fan2.Active.OnValueRemoteUpdate(func(int) { fn() })
	acc.Fan2.RotationSpeed.OnValueRemoteUpdate(func(float64) { fn() })
}
