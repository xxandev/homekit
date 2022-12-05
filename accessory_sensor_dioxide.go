package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorDioxide struct
type AccessorySensorDioxide struct {
	*accessory.A
	CarbonDioxideSensor *service.CarbonDioxideSensor
}

func (acc *AccessorySensorDioxide) GetType() uint8 {
	return acc.A.Type
}

func (acc *AccessorySensorDioxide) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorDioxide) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorDioxide) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorDioxide) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorDioxide) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorDioxide return *SensorDioxide.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessorySensorDioxide(info accessory.Info, args ...interface{}) *AccessorySensorDioxide {
	acc := AccessorySensorDioxide{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.CarbonDioxideSensor = service.NewCarbonDioxideSensor()
	acc.AddS(acc.CarbonDioxideSensor.S)
	return &acc
}

//NewAccSensorDioxide return *SensorDioxide.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccSensorDioxide(id uint64, info accessory.Info, args ...interface{}) *AccessorySensorDioxide {
	acc := AccessorySensorDioxide{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.CarbonDioxideSensor = service.NewCarbonDioxideSensor()
	acc.AddS(acc.CarbonDioxideSensor.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessorySensorDioxide) OnValuesRemoteUpdates(fn func()) {}
