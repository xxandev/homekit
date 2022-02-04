package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryWindow struct
type AccessoryWindow struct {
	*accessory.Accessory
	Window *service.Window
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
