package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorMotion struct
type AccessorySensorMotion struct {
	*accessory.A
	MotionSensor *service.MotionSensor
}

func (acc *AccessorySensorMotion) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorMotion) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorMotion) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorMotion) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorMotion) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorMotion) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorMotion returns *SensorMotion.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessorySensorMotion(info accessory.Info, args ...interface{}) *AccessorySensorMotion {
	acc := AccessorySensorMotion{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.MotionSensor = service.NewMotionSensor()
	acc.AddS(acc.MotionSensor.S)
	return &acc
}

//NewAccSensorMotion returns *SensorMotion.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccSensorMotion(id uint64, info accessory.Info, args ...interface{}) *AccessorySensorMotion {
	acc := AccessorySensorMotion{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.MotionSensor = service.NewMotionSensor()
	acc.AddS(acc.MotionSensor.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessorySensorMotion) OnValuesRemoteUpdates(fn func()) {}
