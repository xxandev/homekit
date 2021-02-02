package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

type serviceSecuritySystemMultifunc struct {
	*service.Service
	SecuritySystemCurrentState *characteristic.SecuritySystemCurrentState
	SecuritySystemTargetState  *characteristic.SecuritySystemTargetState
	SecuritySystemAlarmType    *characteristic.SecuritySystemAlarmType
	StatusFault                *characteristic.StatusFault
	StatusTampered             *characteristic.StatusTampered
}

func newServiceSecuritySystemMultifunc() *serviceSecuritySystemMultifunc {
	svc := serviceSecuritySystemMultifunc{}
	svc.Service = service.New(service.TypeSecuritySystem)

	svc.SecuritySystemCurrentState = characteristic.NewSecuritySystemCurrentState()
	svc.AddCharacteristic(svc.SecuritySystemCurrentState.Characteristic)

	svc.SecuritySystemTargetState = characteristic.NewSecuritySystemTargetState()
	svc.AddCharacteristic(svc.SecuritySystemTargetState.Characteristic)

	svc.SecuritySystemAlarmType = characteristic.NewSecuritySystemAlarmType()
	svc.AddCharacteristic(svc.SecuritySystemAlarmType.Characteristic)

	svc.StatusFault = characteristic.NewStatusFault()
	svc.AddCharacteristic(svc.StatusFault.Characteristic)

	svc.StatusTampered = characteristic.NewStatusTampered()
	svc.AddCharacteristic(svc.StatusTampered.Characteristic)

	return &svc
}

//AccessorySecuritySystemMultifunc struct
type AccessorySecuritySystemMultifunc struct {
	*accessory.Accessory
	SecuritySystemMultifunc *serviceSecuritySystemMultifunc
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
	acc.SecuritySystemMultifunc = newServiceSecuritySystemMultifunc()

	amountArgs := len(args)
	if amountArgs > 0 {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetValue(argToInt(args[0], 0))
	} else {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetValue(0)
	}
	if amountArgs > 1 {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetMinValue(argToInt(args[1], 0))
	} else {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetMinValue(0)
	}
	if amountArgs > 2 {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetMaxValue(argToInt(args[2], 3))
	} else {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetMaxValue(3)
	}
	if amountArgs > 3 {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetStepValue(argToInt(args[3], 1))
	} else {
		acc.SecuritySystemMultifunc.SecuritySystemTargetState.SetStepValue(1)
	}
	acc.AddService(acc.SecuritySystemMultifunc.Service)
	return &acc
}
