package homekit

import (
	"fmt"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryLightbulbSwitch struct
type AccessoryLightbulbSwitch struct {
	*accessory.Accessory
	LightbulbSwitch *service.Lightbulb
}

func (acc *AccessoryLightbulbSwitch) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}
func (acc *AccessoryLightbulbSwitch) GetID() uint64 {
	return acc.Accessory.ID
}
func (acc *AccessoryLightbulbSwitch) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryLightbulbSwitch) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}
func (acc *AccessoryLightbulbSwitch) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryLightbulbSwitch return AccessoryLightbulbSwitch (args... are not used)
func NewAccessoryLightbulbSwitch(info accessory.Info, args ...interface{}) *AccessoryLightbulbSwitch {
	acc := AccessoryLightbulbSwitch{}
	acc.Accessory = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbSwitch = service.NewLightbulb()
	acc.AddService(acc.LightbulbSwitch.Service)
	return &acc
}

func (acc *AccessoryLightbulbSwitch) OnValuesRemoteUpdates(fn func()) {
	acc.LightbulbSwitch.On.OnValueRemoteUpdate(func(bool) { fn() })
}

func (acc *AccessoryLightbulbSwitch) OnValuesRemoteUpdatesPrint() {
	acc.LightbulbSwitch.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%T - %s] remote update on: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
