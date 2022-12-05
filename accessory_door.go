package homekit

import (
	"github.com/brutella/hap/accessory"
	haps "github.com/xxandev/homekit/hap-service"
)

//AccessoryDoor struct
type AccessoryDoor struct {
	*accessory.A
	Door *haps.Door
}

func (acc *AccessoryDoor) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryDoor) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryDoor) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryDoor) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryDoor) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryDoor) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryDoor returns *Door.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args[0](int) - TargetPosition.SetValue(args[0]) default(0)
//  args[1](int) - TargetPosition.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetPosition.SetMaxValue(args[2]) default(100)
//  args[3](int) - TargetPosition.SetStepValue(args[3]) default(1)
func NewAccessoryDoor(info accessory.Info, args ...interface{}) *AccessoryDoor {
	acc := AccessoryDoor{}
	acc.A = accessory.New(info, accessory.TypeDoor)
	acc.Door = haps.NewDoor()
	n := len(args)
	if n > 0 {
		acc.Door.TargetPosition.SetValue(toi(args[0], 0))
	}
	if n > 1 {
		acc.Door.TargetPosition.SetMinValue(toi(args[1], 0))
	}
	if n > 2 {
		acc.Door.TargetPosition.SetMaxValue(toi(args[2], 100))
	}
	if n > 3 {
		acc.Door.TargetPosition.SetStepValue(toi(args[3], 1))
	}
	acc.AddS(acc.Door.S)
	return &acc
}

//NewAccDoor returns *Door.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args[0](int) - TargetPosition.SetValue(args[0]) default(0)
//  args[1](int) - TargetPosition.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetPosition.SetMaxValue(args[2]) default(100)
//  args[3](int) - TargetPosition.SetStepValue(args[3]) default(1)
func NewAccDoor(id uint64, info accessory.Info, args ...interface{}) *AccessoryDoor {
	acc := AccessoryDoor{}
	acc.A = accessory.New(info, accessory.TypeDoor)
	acc.Door = haps.NewDoor()
	n := len(args)
	if n > 0 {
		acc.Door.TargetPosition.SetValue(toi(args[0], 0))
	}
	if n > 1 {
		acc.Door.TargetPosition.SetMinValue(toi(args[1], 0))
	}
	if n > 2 {
		acc.Door.TargetPosition.SetMaxValue(toi(args[2], 100))
	}
	if n > 3 {
		acc.Door.TargetPosition.SetStepValue(toi(args[3], 1))
	}
	acc.AddS(acc.Door.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessoryDoor) OnValuesRemoteUpdates(fn func()) {
	acc.Door.TargetPosition.OnValueRemoteUpdate(func(int) { fn() })
}
