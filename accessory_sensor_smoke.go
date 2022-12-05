package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorSmoke struct
type AccessorySensorSmoke struct {
	*accessory.A
	SmokeSensor *service.SmokeSensor
}

func (acc *AccessorySensorSmoke) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorSmoke) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorSmoke) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorSmoke) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorSmoke) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorSmoke) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorSmoke return *SensorSmoke.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessorySensorSmoke(info accessory.Info, args ...interface{}) *AccessorySensorSmoke {
	acc := AccessorySensorSmoke{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.SmokeSensor = service.NewSmokeSensor()
	acc.AddS(acc.SmokeSensor.S)
	return &acc
}

//NewAccSensorSmoke return *SensorSmoke.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccSensorSmoke(id uint64, info accessory.Info, args ...interface{}) *AccessorySensorSmoke {
	acc := AccessorySensorSmoke{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.SmokeSensor = service.NewSmokeSensor()
	acc.AddS(acc.SmokeSensor.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessorySensorSmoke) OnValuesRemoteUpdates(fn func()) {}
