package homekit

import (
	haps "github.com/alpr777/homekit/hap-service"
	"github.com/brutella/hc/accessory"
)

//AccessorySecuritySystemMultifunc struct
type AccessorySecuritySystemMultifunc struct {
	*accessory.Accessory
	SecuritySystemMultifunc *haps.SecuritySystemMultifunc
}

//NewAccessorySecuritySystemMultifunc return AccessorySecuritySystemMultifunc
//  args[0](int) - SecuritySystemTargetState.SetValue(args[0]) default(0)
//  args[1](int) - SecuritySystemTargetState.SetMinValue(args[1]) default(0)
//  args[2](int) - SecuritySystemTargetState.SetMaxValue(args[2]) default(3)
//  args[3](int) - SecuritySystemTargetState.SetStepValue(args[3]) default(1)
func NewAccessorySecuritySystemMultifunc(info accessory.Info, args ...interface{}) *AccessorySecuritySystemMultifunc {
	acc := AccessorySecuritySystemMultifunc{}
	acc.Accessory = accessory.New(info, accessory.TypeSecuritySystem)
	acc.SecuritySystemMultifunc = haps.NewSecuritySystemMultifunc()

	n := len(args)
	if n > 0 {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetValue(toInt(args[0], 0))
	}
	if n > 1 {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetMinValue(toInt(args[1], 0))
	}
	if n > 2 {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetMaxValue(toInt(args[2], 3))
	}
	if n > 3 {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetStepValue(toInt(args[3], 1))
	}
	acc.AddService(acc.SecuritySystemMultifunc.Service)
	return &acc
}
