package homekit

import (
	"fmt"
	"time"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryFanSwitch struct
type AccessoryFanSwitch struct {
	*accessory.A
	Fan *service.Fan
}

func (acc *AccessoryFanSwitch) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryFanSwitch) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryFanSwitch) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryFanSwitch) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryFanSwitch) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryFanSwitch) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryFanSwitch return AccessoryFanSwitch (args... are not used)
func NewAccessoryFanSwitch(info accessory.Info, args ...interface{}) *AccessoryFanSwitch {
	acc := AccessoryFanSwitch{}
	acc.A = accessory.New(info, accessory.TypeFan)
	acc.Fan = service.NewFan()
	acc.AddS(acc.Fan.S)
	return &acc
}

func (acc *AccessoryFanSwitch) OnValuesRemoteUpdates(fn func()) {
	acc.Fan.On.OnValueRemoteUpdate(func(bool) { fn() })
}

func (acc *AccessoryFanSwitch) OnExample() {
	go func() {
		for range time.Tick(30 * time.Second) {
			acc.Fan.On.SetValue(!acc.Fan.On.Value())
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update on: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.Fan.On.Value())
		}
	}()
	acc.Fan.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update on: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
}
