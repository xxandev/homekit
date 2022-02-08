package homekit

import (
	"fmt"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryWindow struct
type AccessoryWindow struct {
	*accessory.Accessory
	Window *service.Window
}

func (acc *AccessoryWindow) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryWindow) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryWindow) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryWindow) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryWindow) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryWindow returns AccessoryWindow
//  args[0](int) - TargetPosition.SetValue(args[0]) default(0)
//  args[1](int) - TargetPosition.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetPosition.SetMaxValue(args[2]) default(100)
//  args[3](int) - TargetPosition.SetStepValue(args[3]) default(1)
func NewAccessoryWindow(info accessory.Info, args ...interface{}) *AccessoryWindow {
	acc := AccessoryWindow{}
	acc.Accessory = accessory.New(info, accessory.TypeWindow)
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
	acc.AddService(acc.Window.Service)
	return &acc
}

func (acc *AccessoryWindow) OnValuesRemoteUpdates(fn func()) {
	acc.Window.TargetPosition.OnValueRemoteUpdate(func(int) { fn() })
	acc.Window.PositionState.OnValueRemoteUpdate(func(int) { fn() })
}

func (acc *AccessoryWindow) OnExample() {
	acc.Window.TargetPosition.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update target position: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.Window.PositionState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update position state: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
