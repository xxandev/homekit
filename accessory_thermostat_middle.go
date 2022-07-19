package homekit

import (
	"fmt"

	"github.com/brutella/hc/accessory"
	haps "github.com/xxandev/homekit/hap-service"
)

//AccessoryThermostatMultifunc struct
type AccessoryThermostatMiddle struct {
	*accessory.Accessory
	Thermostat *haps.ThermostatMiddle
}

func (acc *AccessoryThermostatMiddle) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryThermostatMiddle) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryThermostatMiddle) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryThermostatMiddle) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryThermostatMiddle) GetAccessory() *accessory.Accessory {
	return acc.Accessory
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

func (acc *AccessoryThermostatMiddle) OnValuesRemoteUpdates(fn func()) {
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(int) { fn() })
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(float64) { fn() })
	acc.Thermostat.TargetRelativeHumidity.OnValueRemoteUpdate(func(float64) { fn() })
	acc.Thermostat.CoolingThresholdTemperature.OnValueRemoteUpdate(func(float64) { fn() })
	acc.Thermostat.HeatingThresholdTemperature.OnValueRemoteUpdate(func(float64) { fn() })
}

func (acc *AccessoryThermostatMiddle) OnExample() {
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update target state: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%T - %s] remote update target temp: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.Thermostat.TargetRelativeHumidity.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%T - %s] remote update relative humidity: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.Thermostat.CoolingThresholdTemperature.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%T - %s] remote update cooling threshold temp: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.Thermostat.HeatingThresholdTemperature.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%T - %s] remote update heating threshold temp: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
