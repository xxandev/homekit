package homekit

import (
	"fmt"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryLightbulbColored struct
type AccessoryLightbulbColored struct {
	*accessory.Accessory
	LightbulbColored *service.ColoredLightbulb
}

func (acc *AccessoryLightbulbColored) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryLightbulbColored) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryLightbulbColored) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryLightbulbColored) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryLightbulbColored) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryLightbulbColored return AccessoryLightbulbColored (args... are not used)
func NewAccessoryLightbulbColored(info accessory.Info, args ...interface{}) *AccessoryLightbulbColored {
	acc := AccessoryLightbulbColored{}
	acc.Accessory = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbColored = service.NewColoredLightbulb()
	acc.AddService(acc.LightbulbColored.Service)
	return &acc
}

func (acc *AccessoryLightbulbColored) OnValuesRemoteUpdates(fn func()) {
	acc.LightbulbColored.On.OnValueRemoteUpdate(func(bool) { fn() })
	acc.LightbulbColored.Brightness.OnValueRemoteUpdate(func(int) { fn() })
	acc.LightbulbColored.Saturation.OnValueRemoteUpdate(func(float64) { fn() })
	acc.LightbulbColored.Hue.OnValueRemoteUpdate(func(float64) { fn() })
}

func (acc *AccessoryLightbulbColored) OnValuesRemoteUpdatesPrint() {
	acc.LightbulbColored.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%T - %s] remote update on: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.LightbulbColored.Brightness.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update brightness: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.LightbulbColored.Saturation.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%T - %s] remote update saturation: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.LightbulbColored.Hue.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%T - %s] remote update hue: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
