package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorLight struct
type AccessorySensorLight struct {
	*accessory.A
	LightSensor *service.LightSensor
}

func (acc *AccessorySensorLight) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorLight) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorLight) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorLight) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorLight) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorLight) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorLight returns *SensorLight.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessorySensorLight(info accessory.Info, args ...interface{}) *AccessorySensorLight {
	acc := AccessorySensorLight{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.LightSensor = service.NewLightSensor()
	acc.AddS(acc.LightSensor.S)
	return &acc
}

//NewAccSensorLight returns *SensorLight.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccSensorLight(id uint64, info accessory.Info, args ...interface{}) *AccessorySensorLight {
	acc := AccessorySensorLight{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.LightSensor = service.NewLightSensor()
	acc.AddS(acc.LightSensor.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessorySensorLight) OnValuesRemoteUpdates(fn func()) {}
