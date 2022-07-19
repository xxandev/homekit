package homekit

import (
	"fmt"
	"time"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorAirQuality struct
type AccessorySensorAirQuality struct {
	*accessory.A
	AirQualitySensor *service.AirQualitySensor
}

func (acc *AccessorySensorAirQuality) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorAirQuality) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorAirQuality) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorAirQuality) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorAirQuality) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorAirQuality) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorAirQuality return AccessorySensorAirQuality args... (args... are not used)
func NewAccessorySensorAirQuality(info accessory.Info, args ...interface{}) *AccessorySensorAirQuality {
	acc := AccessorySensorAirQuality{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.AirQualitySensor = service.NewAirQualitySensor()
	acc.AddS(acc.AirQualitySensor.S)
	return &acc
}

func (acc *AccessorySensorAirQuality) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorAirQuality) OnExample() {
	go func() {
		for range time.Tick(10 * time.Second) {
			if acc.AirQualitySensor.AirQuality.Value() >= 5 {
				acc.AirQualitySensor.AirQuality.SetValue(0)
			} else {
				acc.AirQualitySensor.AirQuality.SetValue(acc.AirQualitySensor.AirQuality.Value() + 1)
			}
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update current state: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.AirQualitySensor.AirQuality.Value())
		}
	}()
}
