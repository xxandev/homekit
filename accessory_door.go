package homekit

import (
	"fmt"

	haps "github.com/alpr777/homekit/hap-service"
	"github.com/brutella/hc/accessory"
)

//AccessoryDoor struct
type AccessoryDoor struct {
	*accessory.Accessory
	Door *haps.Door
}

func (acc *AccessoryDoor) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryDoor) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryDoor) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryDoor) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryDoor) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryDoor returns AccessoryDoor
//  args[0](int) - TargetPosition.SetValue(args[0]) default(0)
//  args[1](int) - TargetPosition.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetPosition.SetMaxValue(args[2]) default(100)
//  args[3](int) - TargetPosition.SetStepValue(args[3]) default(1)
func NewAccessoryDoor(info accessory.Info, args ...interface{}) *AccessoryDoor {
	acc := AccessoryDoor{}
	acc.Accessory = accessory.New(info, accessory.TypeDoor)
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
	acc.AddService(acc.Door.Service)
	return &acc
}

func (acc *AccessoryDoor) OnValuesRemoteUpdates(fn func()) {
	acc.Door.TargetPosition.OnValueRemoteUpdate(func(int) { fn() })
}

func (acc *AccessoryDoor) OnValuesRemoteUpdatesPrint() {
	acc.Door.TargetPosition.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update target position: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
