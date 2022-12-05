package homekit

import (
	"github.com/brutella/hap/accessory"
	haps "github.com/xxandev/homekit/hap-service"
)

//AccessoryLightbulbDimmer struct
type AccessoryLightbulbDimmer struct {
	*accessory.A
	LightbulbDimmer *haps.LightbulbDimmer
}

func (acc *AccessoryLightbulbDimmer) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryLightbulbDimmer) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryLightbulbDimmer) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryLightbulbDimmer) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryLightbulbDimmer) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryLightbulbDimmer) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryLightbulbDimmer return *LightbulbDimmer.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessoryLightbulbDimmer(info accessory.Info, args ...interface{}) *AccessoryLightbulbDimmer {
	acc := AccessoryLightbulbDimmer{}
	acc.A = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbDimmer = haps.NewLightbulbDimmer()
	acc.AddS(acc.LightbulbDimmer.S)
	return &acc
}

//NewAccLightbulbDimmer return *LightbulbDimmer.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccLightbulbDimmer(id uint64, info accessory.Info, args ...interface{}) *AccessoryLightbulbDimmer {
	acc := AccessoryLightbulbDimmer{}
	acc.A = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbDimmer = haps.NewLightbulbDimmer()
	acc.AddS(acc.LightbulbDimmer.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessoryLightbulbDimmer) OnValuesRemoteUpdates(fn func()) {
	acc.LightbulbDimmer.On.OnValueRemoteUpdate(func(bool) { fn() })
	acc.LightbulbDimmer.Brightness.OnValueRemoteUpdate(func(int) { fn() })
}
