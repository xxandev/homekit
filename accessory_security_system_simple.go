package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySecuritySystemSimple struct
type AccessorySecuritySystemSimple struct {
	*accessory.Accessory
	SecuritySystemSimple *service.SecuritySystem
}

//NewAccessorySecuritySystemSimple returns AccessorySecuritySystemSimple
//
//args[0](int) - SecuritySystemTargetState.SetValue(args[0]) default(0)
//
//args[1](int) - SecuritySystemTargetState.SetMinValue(args[1]) default(0)
//
//args[2](int) - SecuritySystemTargetState.SetMaxValue(args[2]) default(3)
//
//args[3](int) - SecuritySystemTargetState.SetStepValue(args[3]) default(1)
func NewAccessorySecuritySystemSimple(info accessory.Info, args ...interface{}) *AccessorySecuritySystemSimple {
	acc := AccessorySecuritySystemSimple{}
	acc.Accessory = accessory.New(info, accessory.TypeSecuritySystem)
	acc.SecuritySystemSimple = service.NewSecuritySystem()
	amountArgs := len(args)
	if amountArgs > 0 {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetValue(argToInt(args[0], 0))
	} else {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetValue(0)
	}
	if amountArgs > 1 {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetMinValue(argToInt(args[1], 0))
	} else {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetMinValue(0)
	}
	if amountArgs > 2 {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetMaxValue(argToInt(args[2], 3))
	} else {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetMaxValue(3)
	}
	if amountArgs > 3 {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetStepValue(argToInt(args[3], 1))
	} else {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetStepValue(1)
	}
	acc.AddService(acc.SecuritySystemSimple.Service)
	return &acc
}
