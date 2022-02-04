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
//  args[0](int) - TargetHeatingCoolingState.SetValue(args[0]) default(0)
//  args[1](int) - TargetHeatingCoolingState.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetHeatingCoolingState.SetMaxValue(args[2]) default(3)
//  args[3](int) - TargetHeatingCoolingState.SetStepValue(args[3]) default(1)
//  args[4](float64) - TargetTemperature.SetValue(args[0]) default(25.0)
//  args[5](float64) - TargetTemperature.SetMinValue(args[1]) default(10.0)
//  args[6](float64) - TargetTemperature.SetMaxValue(args[2]) default(40.0)
//  args[7](float64) - TargetTemperature.SetStepValue(args[3]) default(1.0)
func NewAccessoryThermostat(info accessory.Info, args ...interface{}) *AccessoryThermostat {
	acc := AccessoryThermostat{}
	acc.Accessory = accessory.New(info, accessory.TypeThermostat)
	acc.Thermostat = service.NewThermostat()

	n := len(args)
	if n > 0 {
		acc.Thermostat.TargetHeatingCoolingState.SetValue(toi(args[0], 0))
	}
	if n > 1 {
		acc.Thermostat.TargetHeatingCoolingState.SetMinValue(toi(args[1], 0))
	}
	if n > 2 {
		acc.Thermostat.TargetHeatingCoolingState.SetMaxValue(toi(args[2], 3))
	}
	if n > 3 {
		acc.Thermostat.TargetHeatingCoolingState.SetStepValue(toi(args[3], 1))
	}

	if n > 4 {
		acc.Thermostat.TargetTemperature.SetValue(tof64(args[4], 25.0))
	} else {
		acc.Thermostat.TargetTemperature.SetValue(25.0)
	}
	if n > 5 {
		acc.Thermostat.TargetTemperature.SetMinValue(tof64(args[5], 10.0))
	} else {
		acc.Thermostat.TargetTemperature.SetMinValue(10.0)
	}
	if n > 6 {
		acc.Thermostat.TargetTemperature.SetMaxValue(tof64(args[6], 40.0))
	} else {
		acc.Thermostat.TargetTemperature.SetMaxValue(40.0)
	}
	if n > 7 {
		acc.Thermostat.TargetTemperature.SetStepValue(tof64(args[7], 1.0))
	} else {
		acc.Thermostat.TargetTemperature.SetStepValue(1.0)
	}

	acc.AddService(acc.Thermostat.Service)

	return &acc
}

func (acc *AccessoryThermostat) OnValuesRemoteUpdates(fn func()) {
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(int) { fn() })
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(float64) { fn() })
}
