package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySensorLeak struct
type AccessorySensorLeak struct {
	*accessory.Accessory
	LeakSensor *service.LeakSensor
}

func (acc *AccessorySensorLeak) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessorySensorLeak) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessorySensorLeak) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessorySensorLeak) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessorySensorLeak) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessorySensorLeak return AccessorySensorLeak (args... are not used)
func NewAccessorySensorLeak(info accessory.Info, args ...interface{}) *AccessorySensorLeak {
	acc := AccessorySensorLeak{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.LeakSensor = service.NewLeakSensor()
	acc.AddService(acc.LeakSensor.Service)
	return &acc
}

func (acc *AccessorySensorLeak) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorLeak) OnValuesRemoteUpdatesPrint()     {}
