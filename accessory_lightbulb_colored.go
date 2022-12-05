package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryLightbulbColored struct
type AccessoryLightbulbColored struct {
	*accessory.A
	LightbulbColored *service.ColoredLightbulb
}

func (acc *AccessoryLightbulbColored) GetType() uint8 {
	return uint8(acc.A.Type)
}

func (acc *AccessoryLightbulbColored) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryLightbulbColored) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryLightbulbColored) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryLightbulbColored) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryLightbulbColored) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryLightbulbColored return *LightbulbColored.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessoryLightbulbColored(info accessory.Info, args ...interface{}) *AccessoryLightbulbColored {
	acc := AccessoryLightbulbColored{}
	acc.A = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbColored = service.NewColoredLightbulb()
	acc.AddS(acc.LightbulbColored.S)
	return &acc
}

//NewAccLightbulbColored return *LightbulbColored.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccLightbulbColored(id uint64, info accessory.Info, args ...interface{}) *AccessoryLightbulbColored {
	acc := AccessoryLightbulbColored{}
	acc.A = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbColored = service.NewColoredLightbulb()
	acc.AddS(acc.LightbulbColored.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessoryLightbulbColored) OnValuesRemoteUpdates(fn func()) {
	acc.LightbulbColored.On.OnValueRemoteUpdate(func(bool) { fn() })
	acc.LightbulbColored.Brightness.OnValueRemoteUpdate(func(int) { fn() })
	acc.LightbulbColored.Saturation.OnValueRemoteUpdate(func(float64) { fn() })
	acc.LightbulbColored.Hue.OnValueRemoteUpdate(func(float64) { fn() })
}
