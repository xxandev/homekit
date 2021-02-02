package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

type serviceFanSimple struct {
	*service.Service
	On            *characteristic.On
	RotationSpeed *characteristic.RotationSpeed
}

func newServiceFanSimple() *serviceFanSimple {
	svc := serviceFanSimple{}
	svc.Service = service.New(service.TypeFan)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddCharacteristic(svc.RotationSpeed.Characteristic)

	return &svc
}

//AccessoryFanControlled struct
type AccessoryFanControlled struct {
	*accessory.Accessory
	Fan *serviceFanSimple
}

//NewAccessoryFanControlled return AccessoryFanControlled
//
//args[0](float64) - RotationSpeed.SetValue(args[0]) default(0)
//
//args[1](float64) - RotationSpeed.SetMinValue(args[1]) default(0)
//
//args[2](float64) - RotationSpeed.SetMaxValue(args[2]) default(100)
//
//args[3](float64) - RotationSpeed.SetStepValue(args[3]) default(1)
func NewAccessoryFanControlled(info accessory.Info, args ...interface{}) *AccessoryFanControlled {
	acc := AccessoryFanControlled{}
	acc.Accessory = accessory.New(info, accessory.TypeFan)
	acc.Fan = newServiceFanSimple()
	amountArgs := len(args)
	if amountArgs > 0 {
		acc.Fan.RotationSpeed.SetValue(argToFloat64(args[0], 0.0))
	} else {
		acc.Fan.RotationSpeed.SetValue(0.0)
	}
	if amountArgs > 1 {
		acc.Fan.RotationSpeed.SetMinValue(argToFloat64(args[1], 0.0))
	} else {
		acc.Fan.RotationSpeed.SetMinValue(0.0)
	}
	if amountArgs > 2 {
		acc.Fan.RotationSpeed.SetMaxValue(argToFloat64(args[2], 100.0))
	} else {
		acc.Fan.RotationSpeed.SetMaxValue(100.0)
	}
	if amountArgs > 3 {
		acc.Fan.RotationSpeed.SetStepValue(argToFloat64(args[3], 1.0))
	} else {
		acc.Fan.RotationSpeed.SetStepValue(1.0)
	}
	acc.AddService(acc.Fan.Service)
	return &acc
}
