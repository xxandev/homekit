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

func (acc *AccessorySensorHumidity) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessorySensorHumidity) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessorySensorHumidity) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessorySensorHumidity) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessorySensorHumidity) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessorySensorHumidity returns AccessorySensorHumidity (args... are not used)
func NewAccessorySensorHumidity(info accessory.Info, args ...interface{}) *AccessorySensorHumidity {
	acc := AccessorySensorHumidity{}
	acc.Accessory = accessory.New(info, accessory.TypeThermostat)
	acc.HumiditySensor = service.NewHumiditySensor()
	acc.AddService(acc.HumiditySensor.Service)
	return &acc
}

func (acc *AccessorySensorHumidity) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorHumidity) OnExample()                      {}
