package homekit

import (
	"github.com/alpr777/homekit/hapservices"
	"github.com/brutella/hc/accessory"
)

//AccessorySecuritySystemMultifunc struct
type AccessorySecuritySystemMultifunc struct {
	*accessory.Accessory
	SecuritySystemMultifunc *hapservices.SecuritySystemMultifunc
}

//NewAccessorySecuritySystemMultifunc return AccessorySecuritySystemMultifunc
//
//args[0](int) - SecuritySystemTargetState.SetValue(args[0]) default(0)
//
//args[1](int) - SecuritySystemTargetState.SetMinValue(args[1]) default(0)
//
//args[2](int) - SecuritySystemTargetState.SetMaxValue(args[2]) default(3)
//
//args[3](int) - SecuritySystemTargetState.SetStepValue(args[3]) default(1)
func NewAccessorySecuritySystemMultifunc(info accessory.Info, args ...interface{}) *AccessorySecuritySystemMultifunc {
	acc := AccessorySecuritySystemMultifunc{}
	acc.Accessory = accessory.New(info, accessory.TypeSecuritySystem)
	acc.SecuritySystemMultifunc = hapservices.NewSecuritySystemMultifunc()

	amountArgs := len(args)
	if amountArgs > 0 {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetValue(argToInt(args[0], 0))
	}
	if amountArgs > 1 {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetMinValue(argToInt(args[1], 0))
	}
	if amountArgs > 2 {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetMaxValue(argToInt(args[2], 3))
	}
	if amountArgs > 3 {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetStepValue(argToInt(args[3], 1))
	}
	acc.AddService(acc.SecuritySystemMultifunc.Service)
	return &acc
}
