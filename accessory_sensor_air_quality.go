package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorAirQuality struct
type AccessorySensorAirQuality struct {
	*accessory.A
	AirQualitySensor *service.AirQualitySensor
}

func (acc *AccessorySensorAirQuality) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorAirQuality) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorAirQuality) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorAirQuality) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorAirQuality) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorAirQuality) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorAirQuality returns *SensorAirQuality.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessorySensorAirQuality(info accessory.Info, args ...interface{}) *AccessorySensorAirQuality {
	acc := AccessorySensorAirQuality{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.AirQualitySensor = service.NewAirQualitySensor()
	acc.AddS(acc.AirQualitySensor.S)
	return &acc
}

//NewAccSensorAirQuality returns *SensorAirQuality.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccSensorAirQuality(id uint64, info accessory.Info, args ...interface{}) *AccessorySensorAirQuality {
	acc := AccessorySensorAirQuality{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.AirQualitySensor = service.NewAirQualitySensor()
	acc.AddS(acc.AirQualitySensor.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessorySensorAirQuality) OnValuesRemoteUpdates(fn func()) {}
