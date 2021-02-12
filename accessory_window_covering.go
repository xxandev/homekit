package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryWindowCovering struct
type AccessoryWindowCovering struct {
	*accessory.Accessory
	WindowCovering *service.WindowCovering
}

//NewAccessoryWindowCovering returns AccessoryWindowCovering
//
//args[0](int) - TargetPosition.SetValue(args[0]) default(0)
//
//args[1](int) - TargetPosition.SetMinValue(args[1]) default(0)
//
//args[2](int) - TargetPosition.SetMaxValue(args[2]) default(100)
//
//args[3](int) - TargetPosition.SetStepValue(args[3]) default(1)
func NewAccessoryWindowCovering(info accessory.Info, args ...interface{}) *AccessoryWindowCovering {
	acc := AccessoryWindowCovering{}
	acc.Accessory = accessory.New(info, accessory.TypeWindowCovering)
	acc.WindowCovering = service.NewWindowCovering()

	amountArgs := len(args)
	if amountArgs > 0 {
		acc.WindowCovering.TargetPosition.SetValue(argToInt(args[0], 0))
	}
	if amountArgs > 1 {
		acc.WindowCovering.TargetPosition.SetMinValue(argToInt(args[1], 0))
	}
	if amountArgs > 2 {
		acc.WindowCovering.TargetPosition.SetMaxValue(argToInt(args[2], 100))
	}
	if amountArgs > 3 {
		acc.WindowCovering.TargetPosition.SetStepValue(argToInt(args[3], 1))
	}

	acc.AddService(acc.WindowCovering.Service)

	return &acc
}
