package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorContact struct
type AccessorySensorContact struct {
	*accessory.A
	ContactSensor *service.ContactSensor
}

func (acc *AccessorySensorContact) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorContact) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorContact) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorContact) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorContact) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorContact) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorContact return *SensorContact.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessorySensorContact(info accessory.Info, args ...interface{}) *AccessorySensorContact {
	acc := AccessorySensorContact{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.ContactSensor = service.NewContactSensor()
	acc.AddS(acc.ContactSensor.S)
	return &acc
}

//NewAccSensorContact return *SensorContact.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccSensorContact(id uint64, info accessory.Info, args ...interface{}) *AccessorySensorContact {
	acc := AccessorySensorContact{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.ContactSensor = service.NewContactSensor()
	acc.AddS(acc.ContactSensor.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessorySensorContact) OnValuesRemoteUpdates(fn func()) {}
