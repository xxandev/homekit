package homekit

import (
	"fmt"
	"math/rand"
	"time"

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

//NewAccessorySensorTemperature returns AccessorySensorTemperature (args... are not used)
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

func (acc *AccessorySensorTemperature) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorTemperature) OnExample() {
	go func() {
		for range time.Tick(10 * time.Second) {
			acc.TempSensor.CurrentTemperature.SetValue(rand.Float64())
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update current state: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.TempSensor.CurrentTemperature.Value())
		}
	}()
}
