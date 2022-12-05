package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryWindowCovering struct
type AccessoryWindowCovering struct {
	*accessory.A
	WindowCovering *service.WindowCovering
}

func (acc *AccessoryWindowCovering) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryWindowCovering) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryWindowCovering) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryWindowCovering) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryWindowCovering) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryWindowCovering) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryWindowCovering returns *WindowCovering.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args[0](int) - TargetPosition.SetValue(args[0]) default(0)
//  args[1](int) - TargetPosition.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetPosition.SetMaxValue(args[2]) default(100)
//  args[3](int) - TargetPosition.SetStepValue(args[3]) default(1)
func NewAccessoryWindowCovering(info accessory.Info, args ...interface{}) *AccessoryWindowCovering {
	acc := AccessoryWindowCovering{}
	acc.A = accessory.New(info, accessory.TypeWindowCovering)
	acc.WindowCovering = service.NewWindowCovering()

	n := len(args)
	if n > 0 {
		acc.WindowCovering.TargetPosition.SetValue(toi(args[0], 0))
	}
	if n > 1 {
		acc.WindowCovering.TargetPosition.SetMinValue(toi(args[1], 0))
	}
	if n > 2 {
		acc.WindowCovering.TargetPosition.SetMaxValue(toi(args[2], 100))
	}
	if n > 3 {
		acc.WindowCovering.TargetPosition.SetStepValue(toi(args[3], 1))
	}

	acc.AddS(acc.WindowCovering.S)

	return &acc
}

//NewAccWindowCovering returns *WindowCovering.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args[0](int) - TargetPosition.SetValue(args[0]) default(0)
//  args[1](int) - TargetPosition.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetPosition.SetMaxValue(args[2]) default(100)
//  args[3](int) - TargetPosition.SetStepValue(args[3]) default(1)
func NewAccWindowCovering(id uint64, info accessory.Info, args ...interface{}) *AccessoryWindowCovering {
	acc := AccessoryWindowCovering{}
	acc.A = accessory.New(info, accessory.TypeWindowCovering)
	acc.WindowCovering = service.NewWindowCovering()
	n := len(args)
	if n > 0 {
		acc.WindowCovering.TargetPosition.SetValue(toi(args[0], 0))
	}
	if n > 1 {
		acc.WindowCovering.TargetPosition.SetMinValue(toi(args[1], 0))
	}
	if n > 2 {
		acc.WindowCovering.TargetPosition.SetMaxValue(toi(args[2], 100))
	}
	if n > 3 {
		acc.WindowCovering.TargetPosition.SetStepValue(toi(args[3], 1))
	}
	acc.AddS(acc.WindowCovering.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessoryWindowCovering) OnValuesRemoteUpdates(fn func()) {
	acc.WindowCovering.TargetPosition.OnValueRemoteUpdate(func(int) { fn() })
	acc.WindowCovering.PositionState.OnValueRemoteUpdate(func(int) { fn() })
}
