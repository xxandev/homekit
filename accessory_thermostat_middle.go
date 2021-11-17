package homekit

import (
	haps "github.com/alpr777/homekit/hap-service"
	"github.com/brutella/hc/accessory"
)

//AccessoryThermostatMultifunc struct
type AccessoryThermostatMiddle struct {
	*accessory.Accessory
	Thermostat *haps.ThermostatMiddle
}

//NewAccessoryThermostatMiddle returns NewAccessoryThermostatMiddle
//  args[0](int) - TargetHeatingCoolingState.SetValue(args[0]) default(0)
//  args[1](int) - TargetHeatingCoolingState.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetHeatingCoolingState.SetMaxValue(args[2]) default(3)
//  args[3](int) - TargetHeatingCoolingState.SetStepValue(args[3]) default(1)
//  args[4](float64) - TargetTemperature.SetValue(args[0]) default(25.0)
//  args[5](float64) - TargetTemperature.SetMinValue(args[1]) default(10.0)
//  args[6](float64) - TargetTemperature.SetMaxValue(args[2]) default(40.0)
//  args[7](float64) - TargetTemperature.SetStepValue(args[3]) default(1.0)
func NewAccessoryThermostatMiddle(info accessory.Info, args ...interface{}) *AccessoryThermostatMiddle {
	acc := AccessoryThermostatMiddle{}
	acc.Accessory = accessory.New(info, accessory.TypeThermostat)
	acc.Thermostat = haps.NewThermostatMiddle()

	n := len(args)
	if n > 0 {
		acc.Thermostat.TargetHeatingCoolingState.SetValue(toInt(args[0], 0))
	}
	if n > 1 {
		acc.Thermostat.TargetHeatingCoolingState.SetMinValue(toInt(args[1], 0))
	}
	if n > 2 {
		acc.Thermostat.TargetHeatingCoolingState.SetMaxValue(toInt(args[2], 3))
	}
	if n > 3 {
		acc.Thermostat.TargetHeatingCoolingState.SetStepValue(toInt(args[3], 1))
	}

	if n > 4 {
		acc.Thermostat.TargetTemperature.SetValue(toFloat64(args[4], 25.0))
	} else {
		acc.Thermostat.TargetTemperature.SetValue(25.0)
	}
	if n > 5 {
		acc.Thermostat.TargetTemperature.SetMinValue(toFloat64(args[5], 10.0))
	} else {
		acc.Thermostat.TargetTemperature.SetMinValue(10.0)
	}
	if n > 6 {
		acc.Thermostat.TargetTemperature.SetMaxValue(toFloat64(args[6], 40.0))
	} else {
		acc.Thermostat.TargetTemperature.SetMaxValue(40.0)
	}
	if n > 7 {
		acc.Thermostat.TargetTemperature.SetStepValue(toFloat64(args[7], 1.0))
	} else {
		acc.Thermostat.TargetTemperature.SetStepValue(1.0)
	}
	acc.AddService(acc.Thermostat.Service)

	return &acc
}
