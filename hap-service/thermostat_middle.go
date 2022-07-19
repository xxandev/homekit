package haps

import (
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
)

//ThermostatMiddle
//	◈ CurrentHeatingCoolingState
//	◈ TargetHeatingCoolingState
//	◈ CurrentTemperature
//	◈ TargetTemperature
//	◈ TemperatureDisplayUnits
//	◇ CoolingThresholdTemperature
//	◇ HeatingThresholdTemperature
//	◇ CurrentRelativeHumidity
//	◇ TargetRelativeHumidity
type ThermostatMiddle struct {
	*service.S
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

//NewThermostatMiddle return *ThermostatMiddle
func NewThermostatMiddle() *ThermostatMiddle {
	svc := ThermostatMiddle{}
	svc.S = service.New(service.TypeThermostat)

	svc.CurrentHeatingCoolingState = characteristic.NewCurrentHeatingCoolingState()
	svc.AddC(svc.CurrentHeatingCoolingState.C)

	svc.TargetHeatingCoolingState = characteristic.NewTargetHeatingCoolingState()
	svc.AddC(svc.TargetHeatingCoolingState.C)

	svc.CurrentTemperature = characteristic.NewCurrentTemperature()
	svc.AddC(svc.CurrentTemperature.C)

	svc.TargetTemperature = characteristic.NewTargetTemperature()
	svc.AddC(svc.TargetTemperature.C)

	svc.TemperatureDisplayUnits = characteristic.NewTemperatureDisplayUnits()
	svc.AddC(svc.TemperatureDisplayUnits.C)

	svc.CoolingThresholdTemperature = characteristic.NewCoolingThresholdTemperature()
	svc.AddC(svc.CoolingThresholdTemperature.C)

	svc.HeatingThresholdTemperature = characteristic.NewHeatingThresholdTemperature()
	svc.AddC(svc.HeatingThresholdTemperature.C)

	svc.CurrentRelativeHumidity = characteristic.NewCurrentRelativeHumidity()
	svc.AddC(svc.CurrentRelativeHumidity.C)

	svc.TargetRelativeHumidity = characteristic.NewTargetRelativeHumidity()
	svc.AddC(svc.TargetRelativeHumidity.C)

	return &svc
}
