package homekit

import (
	"fmt"

	haps "github.com/alpr777/homekit/hap-service"
	"github.com/brutella/hc/accessory"
)

//AccessoryFanControlled struct
type AccessoryFanSpeed struct {
	*accessory.Accessory
	Fan *haps.FanRS
}

func (acc *AccessoryFanSpeed) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryFanSpeed) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryFanSpeed) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryFanSpeed) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryFanSpeed) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryFanSpeed return AccessoryFanSpeed
//  args[0](float64) - RotationSpeed.SetValue(args[0]) default(0)
//  args[1](float64) - RotationSpeed.SetMinValue(args[1]) default(0)
//  args[2](float64) - RotationSpeed.SetMaxValue(args[2]) default(100)
//  args[3](float64) - RotationSpeed.SetStepValue(args[3]) default(1)
func NewAccessoryFanSpeed(info accessory.Info, args ...interface{}) *AccessoryFanSpeed {
	acc := AccessoryFanSpeed{}
	acc.Accessory = accessory.New(info, accessory.TypeFan)
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
	acc.AddService(acc.Fan.Service)
	return &acc
}

func (acc *AccessoryFanSpeed) OnValuesRemoteUpdates(fn func()) {
	acc.Fan.On.OnValueRemoteUpdate(func(bool) { fn() })
	acc.Fan.RotationSpeed.OnValueRemoteUpdate(func(float64) { fn() })
}

func (acc *AccessoryFanSpeed) OnExample() {
	acc.Fan.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%T - %s] remote update on: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.Fan.RotationSpeed.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%T - %s] remote update rotation speed: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
