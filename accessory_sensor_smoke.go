package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySensorSmoke struct
type AccessorySensorSmoke struct {
	*accessory.Accessory
	SmokeSensor *service.SmokeSensor
}

func (acc *AccessorySensorSmoke) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessorySensorSmoke) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessorySensorSmoke) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessorySensorSmoke) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessorySensorSmoke) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessorySensorSmoke return AccessorySensorSmoke (args... are not used)
func NewAccessorySensorSmoke(info accessory.Info, args ...interface{}) *AccessorySensorSmoke {
	acc := AccessorySensorSmoke{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.SmokeSensor = service.NewSmokeSensor()
	acc.AddService(acc.SmokeSensor.Service)
	return &acc
}

func (acc *AccessorySensorSmoke) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorSmoke) OnExample()                      {}
