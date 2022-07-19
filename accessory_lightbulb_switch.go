package homekit

import (
	"fmt"
	"time"

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

//NewAccessoryLightbulbSwitch return AccessoryLightbulbSwitch (args... are not used)
func NewAccessoryLightbulbSwitch(info accessory.Info, args ...interface{}) *AccessoryLightbulbSwitch {
	acc := AccessoryLightbulbSwitch{}
	acc.A = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbSwitch = service.NewLightbulb()
	acc.AddS(acc.LightbulbSwitch.S)
	return &acc
}

func (acc *AccessoryLightbulbSwitch) OnValuesRemoteUpdates(fn func()) {
	acc.LightbulbSwitch.On.OnValueRemoteUpdate(func(bool) { fn() })
}

func (acc *AccessoryLightbulbSwitch) OnExample() {
	go func() {
		for range time.Tick(30 * time.Second) {
			acc.LightbulbSwitch.On.SetValue(!acc.LightbulbSwitch.On.Value())
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update on: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.LightbulbSwitch.On.Value())
		}
	}()
	acc.LightbulbSwitch.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update on: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
}
