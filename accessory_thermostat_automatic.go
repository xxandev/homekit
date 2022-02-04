package homekit

import (
	"fmt"

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

func (acc *AccessoryThermostatAutomatic) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryThermostatAutomatic) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryThermostatAutomatic) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryThermostatAutomatic) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryThermostatAutomatic) GetAccessory() *accessory.Accessory {
	return acc.Accessory
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
		acc.Thermostat.TargetTemperature.SetValue(tof64(args[0], 25.0))
	} else {
		acc.Thermostat.TargetTemperature.SetValue(25.0)
	}
	if n > 1 {
		acc.Thermostat.TargetTemperature.SetMinValue(tof64(args[1], 10.0))
	} else {
		acc.Thermostat.TargetTemperature.SetMinValue(10.0)
	}
	if n > 2 {
		acc.Thermostat.TargetTemperature.SetMaxValue(tof64(args[2], 40.0))
	} else {
		acc.Thermostat.TargetTemperature.SetMaxValue(40.0)
	}
	if n > 3 {
		acc.Thermostat.TargetTemperature.SetStepValue(tof64(args[3], 1.0))
	} else {
		acc.Thermostat.TargetTemperature.SetStepValue(1.0)
	}

	acc.AddService(acc.Thermostat.Service)
	acc.AddService(acc.Switch.Service)
	return &acc
}

func (acc *AccessoryThermostatAutomatic) OnValuesRemoteUpdates(fn func()) {
	acc.Switch.On.OnValueRemoteUpdate(func(bool) { fn() })
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(int) { fn() })
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(float64) { fn() })
}

func (acc *AccessoryThermostatAutomatic) OnValuesRemoteUpdatesPrint() {
	acc.Switch.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%T - %s] remote update on: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update target state: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%T - %s] remote update target temp: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
