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

func (acc *AccessorySensorMonoxide) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessorySensorMonoxide) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessorySensorMonoxide) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessorySensorMonoxide) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessorySensorMonoxide) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessorySensorMonoxide return AccessorySensorMonoxide (args... are not used)
func NewAccessorySensorMonoxide(info accessory.Info, args ...interface{}) *AccessorySensorMonoxide {
	acc := AccessorySensorMonoxide{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.CarbonMonoxideSensor = service.NewCarbonMonoxideSensor()
	acc.AddService(acc.CarbonMonoxideSensor.Service)
	return &acc
}

func (acc *AccessorySensorMonoxide) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorMonoxide) OnValuesRemoteUpdatesPrint()     {}
