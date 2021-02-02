package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySensorMotion struct
type AccessorySensorMotion struct {
	*accessory.Accessory
	MotionSensor *service.MotionSensor
}

//NewAccessorySensorMotion return AccessorySensorMotion (args... are not used)
func NewAccessorySensorMotion(info accessory.Info, args ...interface{}) *AccessorySensorMotion {
	acc := AccessorySensorMotion{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.MotionSensor = service.NewMotionSensor()
	acc.AddService(acc.MotionSensor.Service)
	return &acc
}
