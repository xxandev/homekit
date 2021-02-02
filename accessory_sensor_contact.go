package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySensorContact struct
type AccessorySensorContact struct {
	*accessory.Accessory
	ContactSensor *service.ContactSensor
}

//NewAccessorySensorContact return AccessorySensorAirQuality (args... are not used)
func NewAccessorySensorContact(info accessory.Info, args ...interface{}) *AccessorySensorContact {
	acc := AccessorySensorContact{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.ContactSensor = service.NewContactSensor()
	acc.AddService(acc.ContactSensor.Service)
	return &acc
}
