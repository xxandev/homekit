package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryLightbulbSwitch struct
type AccessoryLightbulbSwitch struct {
	*accessory.A
	LightbulbSwitch *service.Lightbulb
}

func (acc *AccessoryLightbulbSwitch) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryLightbulbSwitch) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryLightbulbSwitch) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryLightbulbSwitch) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryLightbulbSwitch) GetName() string {
	return acc.A.Info.Name.Value()
}
func (acc *AccessoryLightbulbSwitch) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryLightbulbSwitch return *LightbulbSwitch.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessoryLightbulbSwitch(info accessory.Info, args ...interface{}) *AccessoryLightbulbSwitch {
	acc := AccessoryLightbulbSwitch{}
	acc.A = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbSwitch = service.NewLightbulb()
	acc.AddS(acc.LightbulbSwitch.S)
	return &acc
}

//NewAccLightbulbSwitch return *LightbulbSwitch.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccLightbulbSwitch(id uint64, info accessory.Info, args ...interface{}) *AccessoryLightbulbSwitch {
	acc := AccessoryLightbulbSwitch{}
	acc.A = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbSwitch = service.NewLightbulb()
	acc.AddS(acc.LightbulbSwitch.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessoryLightbulbSwitch) OnValuesRemoteUpdates(fn func()) {
	acc.LightbulbSwitch.On.OnValueRemoteUpdate(func(bool) { fn() })
}
