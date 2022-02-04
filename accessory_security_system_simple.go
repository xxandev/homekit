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
//  args[0](int) - SecuritySystemTargetState.SetValue(args[0]) default(0)
//  args[1](int) - SecuritySystemTargetState.SetMinValue(args[1]) default(0)
//  args[2](int) - SecuritySystemTargetState.SetMaxValue(args[2]) default(3)
//  args[3](int) - SecuritySystemTargetState.SetStepValue(args[3]) default(1)
func NewAccessorySecuritySystemSimple(info accessory.Info, args ...interface{}) *AccessorySecuritySystemSimple {
	acc := AccessorySecuritySystemSimple{}
	acc.Accessory = accessory.New(info, accessory.TypeSecuritySystem)
	acc.SecuritySystemSimple = service.NewSecuritySystem()
	n := len(args)
	if n > 0 {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetValue(toi(args[0], 0))
	}
	if n > 1 {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetMinValue(toi(args[1], 0))
	}
	if n > 2 {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetMaxValue(toi(args[2], 3))
	}
	if n > 3 {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetStepValue(toi(args[3], 1))
	}
	acc.AddService(acc.SecuritySystemSimple.Service)
	return &acc
}

func (acc *AccessorySecuritySystemSimple) OnValuesRemoteUpdates(fn func()) {
	acc.SecuritySystemSimple.SecuritySystemTargetState.OnValueRemoteUpdate(func(int) { fn() })
}
