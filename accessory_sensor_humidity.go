package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorHumidity struct
type AccessorySensorHumidity struct {
	*accessory.A
	HumiditySensor *service.HumiditySensor
}

func (acc *AccessorySensorHumidity) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorHumidity) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorHumidity) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorHumidity) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorHumidity) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorHumidity) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorHumidity returns *SensorHumidity.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessorySensorHumidity(info accessory.Info, args ...interface{}) *AccessorySensorHumidity {
	acc := AccessorySensorHumidity{}
	acc.A = accessory.New(info, accessory.TypeThermostat)
	acc.HumiditySensor = service.NewHumiditySensor()
	acc.AddS(acc.HumiditySensor.S)
	return &acc
}

//NewAccSensorHumidity returns *SensorHumidity.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccSensorHumidity(id uint64, info accessory.Info, args ...interface{}) *AccessorySensorHumidity {
	acc := AccessorySensorHumidity{}
	acc.A = accessory.New(info, accessory.TypeThermostat)
	acc.HumiditySensor = service.NewHumiditySensor()
	acc.AddS(acc.HumiditySensor.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessorySensorHumidity) OnValuesRemoteUpdates(fn func()) {}
