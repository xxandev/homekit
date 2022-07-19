package homekit

import (
	"fmt"

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

//NewAccessoryDoor returns AccessoryDoor
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

func (acc *AccessoryDoor) OnValuesRemoteUpdates(fn func()) {
	acc.Door.TargetPosition.OnValueRemoteUpdate(func(int) { fn() })
}

func (acc *AccessoryDoor) OnExample() {
	acc.Door.TargetPosition.OnValueRemoteUpdate(func(v int) {
		acc.Door.CurrentPosition.SetValue(v)
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update target position: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
}
