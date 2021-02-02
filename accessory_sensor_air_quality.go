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

//NewAccessorySensorAirQuality return AccessorySensorAirQuality args... (args... are not used)
func NewAccessorySensorAirQuality(info accessory.Info, args ...interface{}) *AccessorySensorAirQuality {
	acc := AccessorySensorAirQuality{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.AirQualitySensor = service.NewAirQualitySensor()
	acc.AddService(acc.AirQualitySensor.Service)
	return &acc
}
