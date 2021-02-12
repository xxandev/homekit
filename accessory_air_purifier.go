package homekit

import (
	"github.com/alpr777/homekit/hapservices"
	"github.com/brutella/hc/accessory"
)

//AccessoryAirPurifier struct
type AccessoryAirPurifier struct {
	*accessory.Accessory
	AirPurifier *hapservices.AirPurifier
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
	acc.AirPurifier = hapservices.NewAirPurifier()
	amountArgs := len(args)
	if amountArgs > 0 {
		acc.AirPurifier.RotationSpeed.SetValue(argToFloat64(args[0], 0.0))
	}
	if amountArgs > 1 {
		acc.AirPurifier.RotationSpeed.SetMinValue(argToFloat64(args[1], 0.0))
	}
	if amountArgs > 2 {
		acc.AirPurifier.RotationSpeed.SetMaxValue(argToFloat64(args[2], 100.0))
	}
	if amountArgs > 3 {
		acc.AirPurifier.RotationSpeed.SetStepValue(argToFloat64(args[3], 1.0))
	}
	acc.AddService(acc.AirPurifier.Service)

	return &acc
}
