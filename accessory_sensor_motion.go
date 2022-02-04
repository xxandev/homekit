package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySensorMotion struct
type AccessorySensorMotion struct {
	*accessory.Accessory
	MotionSensor *service.MotionSensor
}

func (acc *AccessorySensorMotion) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessorySensorMotion) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessorySensorMotion) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessorySensorMotion) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessorySensorMotion) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessorySensorMotion return AccessorySensorMotion (args... are not used)
func NewAccessorySensorMotion(info accessory.Info, args ...interface{}) *AccessorySensorMotion {
	acc := AccessorySensorMotion{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.MotionSensor = service.NewMotionSensor()
	acc.AddService(acc.MotionSensor.Service)
	return &acc
}

func (acc *AccessorySensorMotion) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorMotion) OnValuesRemoteUpdatesPrint()     {}
