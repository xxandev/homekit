package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySensorLeak struct
type AccessorySensorLeak struct {
	*accessory.Accessory
	LeakSensor *service.LeakSensor
}

//NewAccessorySensorLeak return AccessorySensorLeak (args... are not used)
func NewAccessorySensorLeak(info accessory.Info, args ...interface{}) *AccessorySensorLeak {
	acc := AccessorySensorLeak{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.LeakSensor = service.NewLeakSensor()
	acc.AddService(acc.LeakSensor.Service)
	return &acc
}
