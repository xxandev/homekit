package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//AccessoryThermostat struct
type AccessoryThermostatAutomatic struct {
	*accessory.Accessory
	Thermostat struct {
		*service.Service
		CurrentHeatingCoolingState *characteristic.CurrentHeatingCoolingState
		TargetHeatingCoolingState  *characteristic.TargetHeatingCoolingState
		CurrentTemperature         *characteristic.CurrentTemperature
		TargetTemperature          *characteristic.TargetTemperature
		TemperatureDisplayUnits    *characteristic.TemperatureDisplayUnits
	}
	Switch struct {
		*service.Service
		On *characteristic.On
	}
}

//NewAccessoryThermostatAutomatic returns AccessoryThermostatAutomatic
//  args[0](float64) - TargetTemperature.SetValue(args[0]) default(25.0)
//  args[1](float64) - TargetTemperature.SetMinValue(args[1]) default(10.0)
//  args[2](float64) - TargetTemperature.SetMaxValue(args[2]) default(40.0)
//  args[3](float64) - TargetTemperature.SetStepValue(args[3]) default(1.0)
func NewAccessoryThermostatAutomatic(info accessory.Info, args ...interface{}) *AccessoryThermostatAutomatic {
	acc := AccessoryThermostatAutomatic{}
	acc.Accessory = accessory.New(info, accessory.TypeThermostat)
	acc.Thermostat.Service = service.New(service.TypeThermostat)
	acc.Switch.Service = service.New(service.TypeSwitch)

	acc.Thermostat.CurrentHeatingCoolingState = characteristic.NewCurrentHeatingCoolingState()
	acc.Thermostat.AddCharacteristic(acc.Thermostat.CurrentHeatingCoolingState.Characteristic)

	acc.Thermostat.TargetHeatingCoolingState = characteristic.NewTargetHeatingCoolingState()
	acc.Thermostat.AddCharacteristic(acc.Thermostat.TargetHeatingCoolingState.Characteristic)

	acc.Thermostat.CurrentTemperature = characteristic.NewCurrentTemperature()
	acc.Thermostat.AddCharacteristic(acc.Thermostat.CurrentTemperature.Characteristic)

	acc.Thermostat.TargetTemperature = characteristic.NewTargetTemperature()
	acc.Thermostat.AddCharacteristic(acc.Thermostat.TargetTemperature.Characteristic)

	acc.Thermostat.TemperatureDisplayUnits = characteristic.NewTemperatureDisplayUnits()
	acc.Thermostat.AddCharacteristic(acc.Thermostat.TemperatureDisplayUnits.Characteristic)

	acc.Switch.On = characteristic.NewOn()
	acc.Switch.AddCharacteristic(acc.Switch.On.Characteristic)

	n := len(args)
	acc.Thermostat.TargetHeatingCoolingState.SetValue(3)
	acc.Thermostat.TargetHeatingCoolingState.SetMinValue(3)
	acc.Thermostat.TargetHeatingCoolingState.SetMaxValue(3)
	acc.Thermostat.TargetHeatingCoolingState.SetStepValue(0)

	if n > 0 {
		acc.Thermostat.TargetTemperature.SetValue(toFloat64(args[0], 25.0))
	} else {
		acc.Thermostat.TargetTemperature.SetValue(25.0)
	}
	if n > 1 {
		acc.Thermostat.TargetTemperature.SetMinValue(toFloat64(args[1], 10.0))
	} else {
		acc.Thermostat.TargetTemperature.SetMinValue(10.0)
	}
	if n > 2 {
		acc.Thermostat.TargetTemperature.SetMaxValue(toFloat64(args[2], 40.0))
	} else {
		acc.Thermostat.TargetTemperature.SetMaxValue(40.0)
	}
	if n > 3 {
		acc.Thermostat.TargetTemperature.SetStepValue(toFloat64(args[3], 1.0))
	} else {
		acc.Thermostat.TargetTemperature.SetStepValue(1.0)
	}

	acc.AddService(acc.Thermostat.Service)
	acc.AddService(acc.Switch.Service)
	return &acc
}

func (acc *AccessoryThermostatAutomatic) OnValuesRemoteUpdates(fn func()) {
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(_ int) { fn() })
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(_ float64) { fn() })
}
