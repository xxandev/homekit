package homekit

import (
	"fmt"

	"github.com/brutella/hap/accessory"
	haps "github.com/xxandev/homekit/hap-service"
)

//AccessoryAirPurifier struct
type AccessoryAirPurifier struct {
	*accessory.A
	AirPurifier *haps.AirPurifier
}

func (acc *AccessoryAirPurifier) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryAirPurifier) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryAirPurifier) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryAirPurifier) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryAirPurifier) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryAirPurifier) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryAirPurifier returns AccessoryAirPurifier
//  args[0](float64) - RotationSpeed.SetValue(args[0]) default(0.0)
//  args[1](float64) - RotationSpeed.SetMinValue(args[1]) default(0.0)
//  args[2](float64) - RotationSpeed.SetMaxValue(args[2]) default(100.0)
//  args[3](float64) - RotationSpeed.SetStepValue(args[3]) default(1.0)
func NewAccessoryAirPurifier(info accessory.Info, args ...interface{}) *AccessoryAirPurifier {
	acc := AccessoryAirPurifier{}
	acc.A = accessory.New(info, accessory.TypeAirPurifier)
	acc.AirPurifier = haps.NewAirPurifier()
	n := len(args)
	if n > 0 {
		acc.AirPurifier.RotationSpeed.SetValue(tof64(args[0], 0.0))
	}
	if n > 1 {
		acc.AirPurifier.RotationSpeed.SetMinValue(tof64(args[1], 0.0))
	}
	if n > 2 {
		acc.AirPurifier.RotationSpeed.SetMaxValue(tof64(args[2], 100.0))
	}
	if n > 3 {
		acc.AirPurifier.RotationSpeed.SetStepValue(tof64(args[3], 1.0))
	}
	acc.AddS(acc.AirPurifier.S)
	return &acc
}

func (acc *AccessoryAirPurifier) OnValuesRemoteUpdates(fn func()) {
	acc.AirPurifier.Active.OnValueRemoteUpdate(func(int) { fn() })
	acc.AirPurifier.TargetAirPurifierState.OnValueRemoteUpdate(func(int) { fn() })
	acc.AirPurifier.RotationSpeed.OnValueRemoteUpdate(func(float64) { fn() })
}

func (acc *AccessoryAirPurifier) OnExample() {
	acc.AirPurifier.Active.OnValueRemoteUpdate(func(v int) {
		if v > 0 {
			acc.AirPurifier.CurrentAirPurifierState.SetValue(2)
		} else {
			acc.AirPurifier.CurrentAirPurifierState.SetValue(0)
		}
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update active: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
	acc.AirPurifier.TargetAirPurifierState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update target state: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
	acc.AirPurifier.RotationSpeed.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update rotation speed: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
}
