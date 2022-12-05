package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorLeak struct
type AccessorySensorLeak struct {
	*accessory.A
	LeakSensor *service.LeakSensor
}

func (acc *AccessorySensorLeak) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorLeak) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorLeak) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorLeak) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorLeak) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorLeak) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorLeak return *SensorLeak.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessorySensorLeak(info accessory.Info, args ...interface{}) *AccessorySensorLeak {
	acc := AccessorySensorLeak{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.LeakSensor = service.NewLeakSensor()
	acc.AddS(acc.LeakSensor.S)
	return &acc
}

//NewAccSensorLeak return *SensorLeak.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccSensorLeak(id uint64, info accessory.Info, args ...interface{}) *AccessorySensorLeak {
	acc := AccessorySensorLeak{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.LeakSensor = service.NewLeakSensor()
	acc.AddS(acc.LeakSensor.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessorySensorLeak) OnValuesRemoteUpdates(fn func()) {}
