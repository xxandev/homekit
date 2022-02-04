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

func (acc *AccessorySensorDioxide) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessorySensorDioxide) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessorySensorDioxide) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessorySensorDioxide) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessorySensorDioxide) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessorySensorLight return AccessorySensorLight (args... are not used)
func NewAccessorySensorDioxide(info accessory.Info, args ...interface{}) *AccessorySensorDioxide {
	acc := AccessorySensorDioxide{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.CarbonDioxideSensor = service.NewCarbonDioxideSensor()
	acc.AddService(acc.CarbonDioxideSensor.Service)
	return &acc
}

func (acc *AccessorySensorDioxide) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorDioxide) OnValuesRemoteUpdatesPrint()     {}
