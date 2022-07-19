package homekit

import (
	"fmt"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryWindow struct
type AccessoryWindow struct {
	*accessory.A
	Window *service.Window
}

func (acc *AccessoryWindow) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryWindow) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryWindow) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryWindow) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryWindow) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryWindow) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryWindow returns AccessoryWindow
//  args[0](int) - TargetPosition.SetValue(args[0]) default(0)
//  args[1](int) - TargetPosition.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetPosition.SetMaxValue(args[2]) default(100)
//  args[3](int) - TargetPosition.SetStepValue(args[3]) default(1)
func NewAccessoryWindow(info accessory.Info, args ...interface{}) *AccessoryWindow {
	acc := AccessoryWindow{}
	acc.A = accessory.New(info, accessory.TypeWindow)
	acc.Window = service.NewWindow()
	n := len(args)
	if n > 0 {
		acc.Window.TargetPosition.SetValue(toi(args[0], 0))
	}
	if n > 1 {
		acc.Window.TargetPosition.SetMinValue(toi(args[1], 0))
	}
	if n > 2 {
		acc.Window.TargetPosition.SetMaxValue(toi(args[2], 100))
	}
	if n > 3 {
		acc.Window.TargetPosition.SetStepValue(toi(args[3], 1))
	}
	acc.AddS(acc.Window.S)
	return &acc
}

func (acc *AccessoryWindow) OnValuesRemoteUpdates(fn func()) {
	acc.Window.TargetPosition.OnValueRemoteUpdate(func(int) { fn() })
	acc.Window.PositionState.OnValueRemoteUpdate(func(int) { fn() })
}

func (acc *AccessoryWindow) OnExample() {
	acc.Window.TargetPosition.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update target position: %T - %v \n", acc, acc.A.Info.SerialNumber.Value(), v, v)
	})
	acc.Window.PositionState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update position state: %T - %v \n", acc, acc.A.Info.SerialNumber.Value(), v, v)
	})
}
