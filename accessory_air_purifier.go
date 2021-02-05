package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

type serviceAirPurifier struct {
	*service.Service
	Active                  *characteristic.Active
	CurrentAirPurifierState *characteristic.CurrentAirPurifierState
	TargetAirPurifierState  *characteristic.TargetAirPurifierState
	RotationSpeed           *characteristic.RotationSpeed
}

func newServiceAirPurifier() *serviceAirPurifier {
	svc := serviceAirPurifier{}
	svc.Service = service.New(service.TypeAirPurifier)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.CurrentAirPurifierState = characteristic.NewCurrentAirPurifierState()
	svc.AddCharacteristic(svc.CurrentAirPurifierState.Characteristic)

	svc.TargetAirPurifierState = characteristic.NewTargetAirPurifierState()
	svc.AddCharacteristic(svc.TargetAirPurifierState.Characteristic)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddCharacteristic(svc.RotationSpeed.Characteristic)

	return &svc
}

//AccessoryAirPurifier struct
type AccessoryAirPurifier struct {
	*accessory.Accessory
	AirPurifier *serviceAirPurifier
}

//NewAccessoryAirPurifier returns AccessoryAirPurifier
//
//args[0](float64) - RotationSpeed.SetValue(args[0]) default(0.0)
//
//args[1](float64) - RotationSpeed.SetMinValue(args[1]) default(0.0)
//
//args[2](float64) - RotationSpeed.SetMaxValue(args[2]) default(100.0)
//
//args[3](float64) - RotationSpeed.SetStepValue(args[3]) default(1.0)
func NewAccessoryAirPurifier(info accessory.Info, args ...interface{}) *AccessoryAirPurifier {
	acc := AccessoryAirPurifier{}
	acc.Accessory = accessory.New(info, accessory.TypeAirPurifier)
	acc.AirPurifier = newServiceAirPurifier()
	amountArgs := len(args)
	if amountArgs > 0 {
		acc.AirPurifier.RotationSpeed.SetValue(argToFloat64(args[0], 0.0))
	} else {
		acc.AirPurifier.RotationSpeed.SetValue(0.0)
	}
	if amountArgs > 1 {
		acc.AirPurifier.RotationSpeed.SetMinValue(argToFloat64(args[1], 0.0))
	} else {
		acc.AirPurifier.RotationSpeed.SetMinValue(0.0)
	}
	if amountArgs > 2 {
		acc.AirPurifier.RotationSpeed.SetMaxValue(argToFloat64(args[2], 100.0))
	} else {
		acc.AirPurifier.RotationSpeed.SetMaxValue(100.00)
	}
	if amountArgs > 3 {
		acc.AirPurifier.RotationSpeed.SetStepValue(argToFloat64(args[3], 1.0))
	} else {
		acc.AirPurifier.RotationSpeed.SetStepValue(1.0)
	}
	acc.AddService(acc.AirPurifier.Service)

	return &acc
}
