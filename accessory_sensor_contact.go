package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySensorContact struct
type AccessorySensorContact struct {
	*accessory.Accessory
	ContactSensor *service.ContactSensor
}

func (acc *AccessorySensorContact) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessorySensorContact) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessorySensorContact) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessorySensorContact) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessorySensorContact) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessorySensorContact return AccessorySensorAirQuality (args... are not used)
func NewAccessorySensorContact(info accessory.Info, args ...interface{}) *AccessorySensorContact {
	acc := AccessorySensorContact{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.ContactSensor = service.NewContactSensor()
	acc.AddService(acc.ContactSensor.Service)
	return &acc
}

func (acc *AccessorySensorContact) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorContact) OnExample()                      {}
