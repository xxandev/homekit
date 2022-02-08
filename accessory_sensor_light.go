package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySensorLight struct
type AccessorySensorLight struct {
	*accessory.Accessory
	LightSensor *service.LightSensor
}

func (acc *AccessorySensorLight) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessorySensorLight) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessorySensorLight) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessorySensorLight) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessorySensorLight) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessorySensorLight return AccessorySensorLight (args... are not used)
func NewAccessorySensorLight(info accessory.Info, args ...interface{}) *AccessorySensorLight {
	acc := AccessorySensorLight{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.LightSensor = service.NewLightSensor()
	acc.AddService(acc.LightSensor.Service)
	return &acc
}

func (acc *AccessorySensorLight) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorLight) OnExample()                      {}
