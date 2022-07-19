package homekit

import (
	"fmt"
	"time"

	"github.com/brutella/hap/accessory"
	haps "github.com/xxandev/homekit/hap-service"
)

//AccessoryFanControlled struct
type AccessoryFanRS struct {
	*accessory.A
	Fan *haps.FanRS
}

func (acc *AccessoryFanRS) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryFanRS) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryFanRS) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryFanRS) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryFanRS) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryFanRS) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryFanRS return AccessoryFanRS
//  args[0](float64) - RotationSpeed.SetValue(args[0]) default(0)
//  args[1](float64) - RotationSpeed.SetMinValue(args[1]) default(0)
//  args[2](float64) - RotationSpeed.SetMaxValue(args[2]) default(100)
//  args[3](float64) - RotationSpeed.SetStepValue(args[3]) default(1)
func NewAccessoryFanRS(info accessory.Info, args ...interface{}) *AccessoryFanRS {
	acc := AccessoryFanRS{}
	acc.A = accessory.New(info, accessory.TypeFan)
	acc.Fan = haps.NewFanRS()
	n := len(args)
	if n > 0 {
		acc.Fan.RotationSpeed.SetValue(tof64(args[0], 0.0))
	}
	if n > 1 {
		acc.Fan.RotationSpeed.SetMinValue(tof64(args[1], 0.0))
	}
	if n > 2 {
		acc.Fan.RotationSpeed.SetMaxValue(tof64(args[2], 100.0))
	}
	if n > 3 {
		acc.Fan.RotationSpeed.SetStepValue(tof64(args[3], 1.0))
	}
	acc.AddS(acc.Fan.S)
	return &acc
}

func (acc *AccessoryFanRS) OnValuesRemoteUpdates(fn func()) {
	acc.Fan.On.OnValueRemoteUpdate(func(bool) { fn() })
	acc.Fan.RotationSpeed.OnValueRemoteUpdate(func(float64) { fn() })
}

func (acc *AccessoryFanRS) OnExample() {
	go func() {
		for range time.Tick(30 * time.Second) {
			acc.Fan.On.SetValue(!acc.Fan.On.Value())
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update on: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.Fan.On.Value())
		}
	}()
	acc.Fan.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update on: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
	acc.Fan.RotationSpeed.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update rotation speed: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
}
