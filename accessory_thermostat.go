package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryThermostat struct
type AccessoryThermostat struct {
	*accessory.Accessory
	Thermostat *service.Thermostat
}

//NewAccessoryThermostat returns AccessoryThermostat
//
//args[0](int) - TargetHeatingCoolingState.SetValue(args[0]) default(0)
//
//args[1](int) - TargetHeatingCoolingState.SetMinValue(args[1]) default(0)
//
//args[2](int) - TargetHeatingCoolingState.SetMaxValue(args[2]) default(3)
//
//args[3](int) - TargetHeatingCoolingState.SetStepValue(args[3]) default(1)
//
//args[4](float64) - TargetTemperature.SetValue(args[0]) default(25.0)
//
//args[5](float64) - TargetTemperature.SetMinValue(args[1]) default(10.0)
//
//args[6](float64) - TargetTemperature.SetMaxValue(args[2]) default(40.0)
//
//args[7](float64) - TargetTemperature.SetStepValue(args[3]) default(1.0)
func NewAccessoryThermostat(info accessory.Info, args ...interface{}) *AccessoryThermostat {
	acc := AccessoryThermostat{}
	acc.Accessory = accessory.New(info, accessory.TypeThermostat)
	acc.Thermostat = service.NewThermostat()

	amountArgs := len(args)
	if amountArgs > 0 {
		acc.Thermostat.TargetHeatingCoolingState.SetValue(argToInt(args[0], 0))
	} else {
		acc.Thermostat.TargetHeatingCoolingState.SetValue(0)
	}
	if amountArgs > 1 {
		acc.Thermostat.TargetHeatingCoolingState.SetMinValue(argToInt(args[1], 0))
	} else {
		acc.Thermostat.TargetHeatingCoolingState.SetMinValue(0)
	}
	if amountArgs > 2 {
		acc.Thermostat.TargetHeatingCoolingState.SetMaxValue(argToInt(args[2], 3))
	} else {
		acc.Thermostat.TargetHeatingCoolingState.SetMaxValue(3)
	}
	if amountArgs > 3 {
		acc.Thermostat.TargetHeatingCoolingState.SetStepValue(argToInt(args[3], 1))
	} else {
		acc.Thermostat.TargetHeatingCoolingState.SetStepValue(1)
	}

	if amountArgs > 4 {
		acc.Thermostat.TargetTemperature.SetValue(argToFloat64(args[4], 25.0))
	} else {
		acc.Thermostat.TargetTemperature.SetValue(25.0)
	}
	if amountArgs > 5 {
		acc.Thermostat.TargetTemperature.SetMinValue(argToFloat64(args[5], 10.0))
	} else {
		acc.Thermostat.TargetTemperature.SetMinValue(10.0)
	}
	if amountArgs > 6 {
		acc.Thermostat.TargetTemperature.SetMaxValue(argToFloat64(args[6], 40.0))
	} else {
		acc.Thermostat.TargetTemperature.SetMaxValue(40.0)
	}
	if amountArgs > 7 {
		acc.Thermostat.TargetTemperature.SetStepValue(argToFloat64(args[7], 1.0))
	} else {
		acc.Thermostat.TargetTemperature.SetStepValue(1.0)
	}

	acc.AddService(acc.Thermostat.Service)

	return &acc
}
