package homekit

import (
	"fmt"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
)

//AccessoryThermostat struct
type AccessoryThermostatAutomatic struct {
	*accessory.A
	Thermostat struct {
		*service.S
		CurrentHeatingCoolingState *characteristic.CurrentHeatingCoolingState
		TargetHeatingCoolingState  *characteristic.TargetHeatingCoolingState
		CurrentTemperature         *characteristic.CurrentTemperature
		TargetTemperature          *characteristic.TargetTemperature
		TemperatureDisplayUnits    *characteristic.TemperatureDisplayUnits
	}
	Switch struct {
		*service.S
		On *characteristic.On
	}
}

func (acc *AccessoryThermostatAutomatic) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryThermostatAutomatic) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryThermostatAutomatic) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryThermostatAutomatic) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryThermostatAutomatic) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryThermostatAutomatic) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryThermostatAutomatic returns AccessoryThermostatAutomatic
//  args[0](float64) - TargetTemperature.SetValue(args[0]) default(25.0)
//  args[1](float64) - TargetTemperature.SetMinValue(args[1]) default(10.0)
//  args[2](float64) - TargetTemperature.SetMaxValue(args[2]) default(40.0)
//  args[3](float64) - TargetTemperature.SetStepValue(args[3]) default(1.0)
func NewAccessoryThermostatAutomatic(info accessory.Info, args ...interface{}) *AccessoryThermostatAutomatic {
	acc := AccessoryThermostatAutomatic{}
	acc.A = accessory.New(info, accessory.TypeThermostat)
	acc.Thermostat.S = service.New(service.TypeThermostat)
	acc.Switch.S = service.New(service.TypeSwitch)

	acc.Thermostat.CurrentHeatingCoolingState = characteristic.NewCurrentHeatingCoolingState()
	acc.Thermostat.AddC(acc.Thermostat.CurrentHeatingCoolingState.C)

	acc.Thermostat.TargetHeatingCoolingState = characteristic.NewTargetHeatingCoolingState()
	acc.Thermostat.AddC(acc.Thermostat.TargetHeatingCoolingState.C)

	acc.Thermostat.CurrentTemperature = characteristic.NewCurrentTemperature()
	acc.Thermostat.AddC(acc.Thermostat.CurrentTemperature.C)

	acc.Thermostat.TargetTemperature = characteristic.NewTargetTemperature()
	acc.Thermostat.AddC(acc.Thermostat.TargetTemperature.C)

	acc.Thermostat.TemperatureDisplayUnits = characteristic.NewTemperatureDisplayUnits()
	acc.Thermostat.AddC(acc.Thermostat.TemperatureDisplayUnits.C)

	acc.Switch.On = characteristic.NewOn()
	acc.Switch.AddC(acc.Switch.On.C)

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

	acc.AddS(acc.Thermostat.S)
	acc.AddS(acc.Switch.S)
	return &acc
}

func (acc *AccessoryThermostatAutomatic) OnValuesRemoteUpdates(fn func()) {
	acc.Switch.On.OnValueRemoteUpdate(func(bool) { fn() })
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(int) { fn() })
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(float64) { fn() })
}

func (acc *AccessoryThermostatAutomatic) OnExample() {
	acc.Switch.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%T - %s] remote update on: %T - %v \n", acc, acc.A.Info.SerialNumber.Value(), v, v)
	})
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update target state: %T - %v \n", acc, acc.A.Info.SerialNumber.Value(), v, v)
	})
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%T - %s] remote update target temp: %T - %v \n", acc, acc.A.Info.SerialNumber.Value(), v, v)
	})
}
