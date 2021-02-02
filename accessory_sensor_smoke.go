package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySensorSmoke struct
type AccessorySensorSmoke struct {
	*accessory.Accessory
	SmokeSensor *service.SmokeSensor
}

//NewAccessorySensorSmoke return AccessorySensorSmoke (args... are not used)
func NewAccessorySensorSmoke(info accessory.Info, args ...interface{}) *AccessorySensorSmoke {
	acc := AccessorySensorSmoke{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.SmokeSensor = service.NewSmokeSensor()
	acc.AddService(acc.SmokeSensor.Service)
	return &acc
}
