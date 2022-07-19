package homekit

import (
	"fmt"

	"github.com/brutella/hap/accessory"
	haps "github.com/xxandev/homekit/hap-service"
)

//AccessoryHumidifierDehumidifier struct
type AccessoryHumidifierDehumidifier struct {
	*accessory.A
	HumidifierDehumidifier *haps.HumidifierDehumidifier
}

func (acc *AccessoryHumidifierDehumidifier) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryHumidifierDehumidifier) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryHumidifierDehumidifier) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryHumidifierDehumidifier) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryHumidifierDehumidifier) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryHumidifierDehumidifier) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryHumidifierDehumidifier returns AccessoryHumidifierDehumidifier
//  args[0](int) - TargetHumidifierDehumidifierState.SetValue(args[0]) default(0)
//  args[1](int) - TargetHumidifierDehumidifierState.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetHumidifierDehumidifierState.SetMaxValue(args[2]) default(2)
//  args[3](int) - TargetHumidifierDehumidifierState.SetStepValue(args[3]) default(1)
func NewAccessoryHumidifierDehumidifier(info accessory.Info, args ...interface{}) *AccessoryHumidifierDehumidifier {
	acc := AccessoryHumidifierDehumidifier{}
	acc.A = accessory.New(info, accessory.TypeDehumidifier)
	acc.HumidifierDehumidifier = haps.NewHumidifierDehumidifier()
	n := len(args)
	if n > 0 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetValue(toi(args[0], 0))
	}
	if n > 1 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMinValue(toi(args[1], 0))
	}
	if n > 2 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMaxValue(toi(args[2], 2))
	}
	if n > 3 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetStepValue(toi(args[3], 1))
	}
	acc.AddS(acc.HumidifierDehumidifier.S)
	return &acc
}

func (acc *AccessoryHumidifierDehumidifier) OnValuesRemoteUpdates(fn func()) {
	acc.HumidifierDehumidifier.Active.OnValueRemoteUpdate(func(int) { fn() })
	acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.OnValueRemoteUpdate(func(int) { fn() })
	acc.HumidifierDehumidifier.RelativeHumidityDehumidifierThreshold.OnValueRemoteUpdate(func(float64) { fn() })
	acc.HumidifierDehumidifier.RelativeHumidityHumidifierThreshold.OnValueRemoteUpdate(func(float64) { fn() })
}

func (acc *AccessoryHumidifierDehumidifier) OnExample() {
	acc.HumidifierDehumidifier.Active.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update active: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
	acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update target state: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
	acc.HumidifierDehumidifier.RelativeHumidityDehumidifierThreshold.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update relative threshold: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
	acc.HumidifierDehumidifier.RelativeHumidityHumidifierThreshold.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update relative threshold: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
}
