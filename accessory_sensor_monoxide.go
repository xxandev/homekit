package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySensorMonoxide struct
type AccessorySensorMonoxide struct {
	*accessory.Accessory
	CarbonMonoxideSensor *service.CarbonMonoxideSensor
}

//NewAccessorySensorMonoxide return AccessorySensorMonoxide (args... are not used)
func NewAccessorySensorMonoxide(info accessory.Info, args ...interface{}) *AccessorySensorMonoxide {
	acc := AccessorySensorMonoxide{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.CarbonMonoxideSensor = service.NewCarbonMonoxideSensor()
	acc.AddService(acc.CarbonMonoxideSensor.Service)
	return &acc
}
