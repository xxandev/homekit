package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySensorHumidity struct
type AccessorySensorHumidity struct {
	*accessory.Accessory
	HumiditySensor *service.HumiditySensor
}

//NewAccessorySensorHumidity returns AccessorySensorHumidity (args... are not used)
func NewAccessorySensorHumidity(info accessory.Info, args ...interface{}) *AccessorySensorHumidity {
	acc := AccessorySensorHumidity{}
	acc.Accessory = accessory.New(info, accessory.TypeThermostat)
	acc.HumiditySensor = service.NewHumiditySensor()
	acc.AddService(acc.HumiditySensor.Service)
	return &acc
}
