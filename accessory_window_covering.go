package homekit

import (
	"fmt"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryWindowCovering struct
type AccessoryWindowCovering struct {
	*accessory.Accessory
	WindowCovering *service.WindowCovering
}

func (acc *AccessoryWindowCovering) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryWindowCovering) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryWindowCovering) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryWindowCovering) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryWindowCovering) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryWindowCovering returns AccessoryWindowCovering
//  args[0](int) - TargetPosition.SetValue(args[0]) default(0)
//  args[1](int) - TargetPosition.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetPosition.SetMaxValue(args[2]) default(100)
//  args[3](int) - TargetPosition.SetStepValue(args[3]) default(1)
func NewAccessoryWindowCovering(info accessory.Info, args ...interface{}) *AccessoryWindowCovering {
	acc := AccessoryWindowCovering{}
	acc.Accessory = accessory.New(info, accessory.TypeWindowCovering)
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

	acc.AddService(acc.WindowCovering.Service)

	return &acc
}

func (acc *AccessoryWindowCovering) OnValuesRemoteUpdates(fn func()) {
	acc.WindowCovering.TargetPosition.OnValueRemoteUpdate(func(int) { fn() })
	acc.WindowCovering.PositionState.OnValueRemoteUpdate(func(int) { fn() })
}

func (acc *AccessoryWindowCovering) OnExample() {
	acc.WindowCovering.TargetPosition.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update target position: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.WindowCovering.PositionState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update position state: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
