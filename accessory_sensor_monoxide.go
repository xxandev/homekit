package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorMonoxide struct
type AccessorySensorMonoxide struct {
	*accessory.A
	CarbonMonoxideSensor *service.CarbonMonoxideSensor
}

func (acc *AccessorySensorMonoxide) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorMonoxide) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorMonoxide) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorMonoxide) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorMonoxide) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorMonoxide) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorMonoxide return *SensorMonoxide.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessorySensorMonoxide(info accessory.Info, args ...interface{}) *AccessorySensorMonoxide {
	acc := AccessorySensorMonoxide{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.CarbonMonoxideSensor = service.NewCarbonMonoxideSensor()
	acc.AddS(acc.CarbonMonoxideSensor.S)
	return &acc
}

//NewAccSensorMonoxide return *SensorMonoxide.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccSensorMonoxide(id uint64, info accessory.Info, args ...interface{}) *AccessorySensorMonoxide {
	acc := AccessorySensorMonoxide{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.CarbonMonoxideSensor = service.NewCarbonMonoxideSensor()
	acc.AddS(acc.CarbonMonoxideSensor.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessorySensorMonoxide) OnValuesRemoteUpdates(fn func()) {}
