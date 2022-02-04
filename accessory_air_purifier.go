package homekit

import (
	"fmt"

	haps "github.com/alpr777/homekit/hap-service"
	"github.com/brutella/hc/accessory"
)

//AccessoryAirPurifier struct
type AccessoryAirPurifier struct {
	*accessory.Accessory
	AirPurifier *haps.AirPurifier
}

func (acc *AccessoryAirPurifier) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryAirPurifier) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryAirPurifier) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryAirPurifier) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryAirPurifier) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryAirPurifier returns AccessoryAirPurifier
//  args[0](float64) - RotationSpeed.SetValue(args[0]) default(0.0)
//  args[1](float64) - RotationSpeed.SetMinValue(args[1]) default(0.0)
//  args[2](float64) - RotationSpeed.SetMaxValue(args[2]) default(100.0)
//  args[3](float64) - RotationSpeed.SetStepValue(args[3]) default(1.0)
func NewAccessoryAirPurifier(info accessory.Info, args ...interface{}) *AccessoryAirPurifier {
	acc := AccessoryAirPurifier{}
	acc.Accessory = accessory.New(info, accessory.TypeAirPurifier)
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
	acc.AddService(acc.AirPurifier.Service)
	return &acc
}

func (acc *AccessoryAirPurifier) OnValuesRemoteUpdates(fn func()) {
	acc.AirPurifier.Active.OnValueRemoteUpdate(func(int) { fn() })
	acc.AirPurifier.TargetAirPurifierState.OnValueRemoteUpdate(func(int) { fn() })
	acc.AirPurifier.RotationSpeed.OnValueRemoteUpdate(func(float64) { fn() })
}

func (acc *AccessoryAirPurifier) OnValuesRemoteUpdatesPrint() {
	acc.AirPurifier.Active.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update active: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.AirPurifier.TargetAirPurifierState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update target state: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.AirPurifier.RotationSpeed.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%T - %s] remote update rotation speed: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
