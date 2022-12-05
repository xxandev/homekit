package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorTemperature struct
type AccessorySensorTemperature struct {
	*accessory.A
	TempSensor *service.TemperatureSensor
}

func (acc *AccessorySensorTemperature) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorTemperature) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorTemperature) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorTemperature) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorTemperature) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorTemperature) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorTemperature returns *SensorTemperature.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessorySensorTemperature(info accessory.Info, args ...interface{}) *AccessorySensorTemperature {
	acc := AccessorySensorTemperature{}
	acc.A = accessory.New(info, accessory.TypeThermostat)
	acc.TempSensor = service.NewTemperatureSensor()
	acc.TempSensor.CurrentTemperature.SetValue(0)
	acc.TempSensor.CurrentTemperature.SetMinValue(-99.00)
	acc.TempSensor.CurrentTemperature.SetMaxValue(99.00)
	acc.TempSensor.CurrentTemperature.SetStepValue(0.1)
	acc.AddS(acc.TempSensor.S)
	return &acc
}

//NewAccessorySensorTemperature returns *SensorTemperature.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccSensorTemperature(id uint64, info accessory.Info, args ...interface{}) *AccessorySensorTemperature {
	acc := AccessorySensorTemperature{}
	acc.A = accessory.New(info, accessory.TypeThermostat)
	acc.TempSensor = service.NewTemperatureSensor()
	acc.TempSensor.CurrentTemperature.SetValue(0)
	acc.TempSensor.CurrentTemperature.SetMinValue(-99.00)
	acc.TempSensor.CurrentTemperature.SetMaxValue(99.00)
	acc.TempSensor.CurrentTemperature.SetStepValue(0.1)
	acc.AddS(acc.TempSensor.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessorySensorTemperature) OnValuesRemoteUpdates(fn func()) {}
