package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySecuritySystem struct
type AccessorySecuritySystem struct {
	*accessory.A
	SecuritySystem *service.SecuritySystem
}

func (acc *AccessorySecuritySystem) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySecuritySystem) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySecuritySystem) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySecuritySystem) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySecuritySystem) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySecuritySystem) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySecuritySystem returns *SecuritySystem.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args[0](int) - SecuritySystemTargetState.SetValue(args[0]) default(0)
//  args[1](int) - SecuritySystemTargetState.SetMinValue(args[1]) default(0)
//  args[2](int) - SecuritySystemTargetState.SetMaxValue(args[2]) default(3)
//  args[3](int) - SecuritySystemTargetState.SetStepValue(args[3]) default(1)
func NewAccessorySecuritySystem(info accessory.Info, args ...interface{}) *AccessorySecuritySystem {
	acc := AccessorySecuritySystem{}
	acc.A = accessory.New(info, accessory.TypeSecuritySystem)
	acc.SecuritySystem = service.NewSecuritySystem()
	n := len(args)
	if n > 0 {
		acc.SecuritySystem.SecuritySystemTargetState.SetValue(toi(args[0], 0))
	}
	if n > 1 {
		acc.SecuritySystem.SecuritySystemTargetState.SetMinValue(toi(args[1], 0))
	}
	if n > 2 {
		acc.SecuritySystem.SecuritySystemTargetState.SetMaxValue(toi(args[2], 3))
	}
	if n > 3 {
		acc.SecuritySystem.SecuritySystemTargetState.SetStepValue(toi(args[3], 1))
	}
	acc.AddS(acc.SecuritySystem.S)
	return &acc
}

//NewAccSecuritySystem returns *SecuritySystem.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args[0](int) - SecuritySystemTargetState.SetValue(args[0]) default(0)
//  args[1](int) - SecuritySystemTargetState.SetMinValue(args[1]) default(0)
//  args[2](int) - SecuritySystemTargetState.SetMaxValue(args[2]) default(3)
//  args[3](int) - SecuritySystemTargetState.SetStepValue(args[3]) default(1)
func NewAccSecuritySystem(id uint64, info accessory.Info, args ...interface{}) *AccessorySecuritySystem {
	acc := AccessorySecuritySystem{}
	acc.A = accessory.New(info, accessory.TypeSecuritySystem)
	acc.SecuritySystem = service.NewSecuritySystem()
	n := len(args)
	if n > 0 {
		acc.SecuritySystem.SecuritySystemTargetState.SetValue(toi(args[0], 0))
	}
	if n > 1 {
		acc.SecuritySystem.SecuritySystemTargetState.SetMinValue(toi(args[1], 0))
	}
	if n > 2 {
		acc.SecuritySystem.SecuritySystemTargetState.SetMaxValue(toi(args[2], 3))
	}
	if n > 3 {
		acc.SecuritySystem.SecuritySystemTargetState.SetStepValue(toi(args[3], 1))
	}
	acc.AddS(acc.SecuritySystem.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessorySecuritySystem) OnValuesRemoteUpdates(fn func()) {
	acc.SecuritySystem.SecuritySystemTargetState.OnValueRemoteUpdate(func(int) { fn() })
}
