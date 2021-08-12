package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySensorDioxide struct
type AccessorySensorDioxide struct {
	*accessory.Accessory
	CarbonDioxideSensor *service.CarbonDioxideSensor
}

//NewAccessorySensorLight return AccessorySensorLight (args... are not used)
func NewAccessorySensorDioxide(info accessory.Info, args ...interface{}) *AccessorySensorDioxide {
	acc := AccessorySensorDioxide{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.CarbonDioxideSensor = service.NewCarbonDioxideSensor()
	acc.AddService(acc.CarbonDioxideSensor.Service)
	return &acc
}
