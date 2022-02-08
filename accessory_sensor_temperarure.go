package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySensorTemperature struct
type AccessorySensorTemperature struct {
	*accessory.Accessory
	TempSensor *service.TemperatureSensor
}

func (acc *AccessorySensorTemperature) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessorySensorTemperature) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessorySensorTemperature) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessorySensorTemperature) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessorySensorTemperature) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessorySensorTemperature returns AccessorySensorTemperature (args... are not used)
func NewAccessorySensorTemperature(info accessory.Info, args ...interface{}) *AccessorySensorTemperature {
	acc := AccessorySensorTemperature{}
	acc.Accessory = accessory.New(info, accessory.TypeThermostat)
	acc.TempSensor = service.NewTemperatureSensor()
	acc.TempSensor.CurrentTemperature.SetValue(0)
	acc.TempSensor.CurrentTemperature.SetMinValue(-99.00)
	acc.TempSensor.CurrentTemperature.SetMaxValue(99.00)
	acc.TempSensor.CurrentTemperature.SetStepValue(0.1)
	acc.AddService(acc.TempSensor.Service)
	return &acc
}

func (acc *AccessorySensorTemperature) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorTemperature) OnExample()                      {}
