package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

type serviceThermostatMultifunc struct {
	*service.Service
	CurrentHeatingCoolingState  *characteristic.CurrentHeatingCoolingState
	TargetHeatingCoolingState   *characteristic.TargetHeatingCoolingState
	CurrentTemperature          *characteristic.CurrentTemperature
	TargetTemperature           *characteristic.TargetTemperature
	TemperatureDisplayUnits     *characteristic.TemperatureDisplayUnits
	CoolingThresholdTemperature *characteristic.CoolingThresholdTemperature
	HeatingThresholdTemperature *characteristic.HeatingThresholdTemperature
	CurrentRelativeHumidity     *characteristic.CurrentRelativeHumidity
	TargetRelativeHumidity      *characteristic.TargetRelativeHumidity
}

func newServiceThermostatMultifunc() *serviceThermostatMultifunc {
	svc := serviceThermostatMultifunc{}
	svc.Service = service.New(service.TypeThermostat)

	svc.CurrentHeatingCoolingState = characteristic.NewCurrentHeatingCoolingState()
	svc.AddCharacteristic(svc.CurrentHeatingCoolingState.Characteristic)

	svc.TargetHeatingCoolingState = characteristic.NewTargetHeatingCoolingState()
	svc.AddCharacteristic(svc.TargetHeatingCoolingState.Characteristic)

	svc.CurrentTemperature = characteristic.NewCurrentTemperature()
	svc.AddCharacteristic(svc.CurrentTemperature.Characteristic)

	svc.TargetTemperature = characteristic.NewTargetTemperature()
	svc.AddCharacteristic(svc.TargetTemperature.Characteristic)

	svc.TemperatureDisplayUnits = characteristic.NewTemperatureDisplayUnits()
	svc.AddCharacteristic(svc.TemperatureDisplayUnits.Characteristic)

	svc.CoolingThresholdTemperature = characteristic.NewCoolingThresholdTemperature()
	svc.AddCharacteristic(svc.CoolingThresholdTemperature.Characteristic)

	svc.HeatingThresholdTemperature = characteristic.NewHeatingThresholdTemperature()
	svc.AddCharacteristic(svc.HeatingThresholdTemperature.Characteristic)

	svc.CurrentRelativeHumidity = characteristic.NewCurrentRelativeHumidity()
	svc.AddCharacteristic(svc.CurrentRelativeHumidity.Characteristic)

	svc.TargetRelativeHumidity = characteristic.NewTargetRelativeHumidity()
	svc.AddCharacteristic(svc.TargetRelativeHumidity.Characteristic)

	return &svc
}

//AccessoryThermostatMultifunc struct
type AccessoryThermostatMultifunc struct {
	*accessory.Accessory
	ThermostatMultifunc *serviceThermostatMultifunc
}

//NewAccessoryThermostatMultifunc returns AccessoryThermostatMultifunc
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
func NewAccessoryThermostatMultifunc(info accessory.Info, args ...interface{}) *AccessoryThermostatMultifunc {
	acc := AccessoryThermostatMultifunc{}
	acc.Accessory = accessory.New(info, accessory.TypeThermostat)
	acc.ThermostatMultifunc = newServiceThermostatMultifunc()
	amountArgs := len(args)
	if amountArgs > 0 {
		acc.ThermostatMultifunc.TargetHeatingCoolingState.SetValue(argToInt(args[0], 0))
	} else {
		acc.ThermostatMultifunc.TargetHeatingCoolingState.SetValue(0)
	}
	if amountArgs > 1 {
		acc.ThermostatMultifunc.TargetHeatingCoolingState.SetMinValue(argToInt(args[1], 0))
	} else {
		acc.ThermostatMultifunc.TargetHeatingCoolingState.SetMinValue(0)
	}
	if amountArgs > 2 {
		acc.ThermostatMultifunc.TargetHeatingCoolingState.SetMaxValue(argToInt(args[2], 3))
	} else {
		acc.ThermostatMultifunc.TargetHeatingCoolingState.SetMaxValue(3)
	}
	if amountArgs > 3 {
		acc.ThermostatMultifunc.TargetHeatingCoolingState.SetStepValue(argToInt(args[3], 1))
	} else {
		acc.ThermostatMultifunc.TargetHeatingCoolingState.SetStepValue(1)
	}

	if amountArgs > 4 {
		acc.ThermostatMultifunc.TargetTemperature.SetValue(argToFloat64(args[4], 25.0))
	} else {
		acc.ThermostatMultifunc.TargetTemperature.SetValue(25.0)
	}
	if amountArgs > 5 {
		acc.ThermostatMultifunc.TargetTemperature.SetMinValue(argToFloat64(args[5], 10.0))
	} else {
		acc.ThermostatMultifunc.TargetTemperature.SetMinValue(10.0)
	}
	if amountArgs > 6 {
		acc.ThermostatMultifunc.TargetTemperature.SetMaxValue(argToFloat64(args[6], 40.0))
	} else {
		acc.ThermostatMultifunc.TargetTemperature.SetMaxValue(40.0)
	}
	if amountArgs > 7 {
		acc.ThermostatMultifunc.TargetTemperature.SetStepValue(argToFloat64(args[7], 1.0))
	} else {
		acc.ThermostatMultifunc.TargetTemperature.SetStepValue(1.0)
	}
	acc.AddService(acc.ThermostatMultifunc.Service)

	return &acc
}
