package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySensorLight struct
type AccessorySensorLight struct {
	*accessory.Accessory
	LightSensor *service.LightSensor
}

//NewAccessorySensorLight return AccessorySensorLight (args... are not used)
func NewAccessorySensorLight(info accessory.Info, args ...interface{}) *AccessorySensorLight {
	acc := AccessorySensorLight{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.LightSensor = service.NewLightSensor()
	acc.AddService(acc.LightSensor.Service)
	return &acc
}
