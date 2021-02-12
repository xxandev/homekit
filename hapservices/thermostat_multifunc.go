package hapservices

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//ThermostatMultifunc (+CurrentHeatingCoolingState, +TargetHeatingCoolingState, +CurrentTemperature, +TargetTemperature,
//+TemperatureDisplayUnits, CoolingThresholdTemperature, HeatingThresholdTemperature, CurrentRelativeHumidity, TargetRelativeHumidity)
type ThermostatMultifunc struct {
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

//NewThermostatMultifunc return *ThermostatMultifunc
func NewThermostatMultifunc() *ThermostatMultifunc {
	svc := ThermostatMultifunc{}
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
