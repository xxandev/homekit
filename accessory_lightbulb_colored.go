package homekit

import (
	"fmt"
	"time"

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

//NewAccessoryLightbulbColored return AccessoryLightbulbColored (args... are not used)
func NewAccessoryLightbulbColored(info accessory.Info, args ...interface{}) *AccessoryLightbulbColored {
	acc := AccessoryLightbulbColored{}
	acc.A = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbColored = service.NewColoredLightbulb()
	acc.AddS(acc.LightbulbColored.S)
	return &acc
}

func (acc *AccessoryLightbulbColored) OnValuesRemoteUpdates(fn func()) {
	acc.LightbulbColored.On.OnValueRemoteUpdate(func(bool) { fn() })
	acc.LightbulbColored.Brightness.OnValueRemoteUpdate(func(int) { fn() })
	acc.LightbulbColored.Saturation.OnValueRemoteUpdate(func(float64) { fn() })
	acc.LightbulbColored.Hue.OnValueRemoteUpdate(func(float64) { fn() })
}

func (acc *AccessoryLightbulbColored) OnExample() {
	go func() {
		for range time.Tick(30 * time.Second) {
			acc.LightbulbColored.On.SetValue(!acc.LightbulbColored.On.Value())
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update on: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.LightbulbColored.On.Value())
		}
	}()
	acc.LightbulbColored.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update on: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
	acc.LightbulbColored.Brightness.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update brightness: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
	acc.LightbulbColored.Saturation.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update saturation: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
	acc.LightbulbColored.Hue.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update hue: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
}
