package homekit

import (
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

//NewAccessoryAirPurifier returns *AirPurifier.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
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

//NewAccAirPurifier returns *AirPurifier.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args[0](float64) - RotationSpeed.SetValue(args[0]) default(0.0)
//  args[1](float64) - RotationSpeed.SetMinValue(args[1]) default(0.0)
//  args[2](float64) - RotationSpeed.SetMaxValue(args[2]) default(100.0)
//  args[3](float64) - RotationSpeed.SetStepValue(args[3]) default(1.0)
func NewAccAirPurifier(id uint64, info accessory.Info, args ...interface{}) *AccessoryAirPurifier {
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
	acc.A.Id = id
	return &acc
}

func (acc *AccessoryAirPurifier) OnValuesRemoteUpdates(fn func()) {
	acc.AirPurifier.Active.OnValueRemoteUpdate(func(int) { fn() })
	acc.AirPurifier.TargetAirPurifierState.OnValueRemoteUpdate(func(int) { fn() })
	acc.AirPurifier.RotationSpeed.OnValueRemoteUpdate(func(float64) { fn() })
}
