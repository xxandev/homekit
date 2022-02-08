package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySensorAirQuality struct
type AccessorySensorAirQuality struct {
	*accessory.Accessory
	AirQualitySensor *service.AirQualitySensor
}

func (acc *AccessorySensorAirQuality) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessorySensorAirQuality) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessorySensorAirQuality) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessorySensorAirQuality) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessorySensorAirQuality) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessorySensorAirQuality return AccessorySensorAirQuality args... (args... are not used)
func NewAccessorySensorAirQuality(info accessory.Info, args ...interface{}) *AccessorySensorAirQuality {
	acc := AccessorySensorAirQuality{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.AirQualitySensor = service.NewAirQualitySensor()
	acc.AddService(acc.AirQualitySensor.Service)
	return &acc
}

func (acc *AccessorySensorAirQuality) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorAirQuality) OnExample()                      {}
