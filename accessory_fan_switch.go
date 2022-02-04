package homekit

import (
	"fmt"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryFanSwitch struct
type AccessoryFanSwitch struct {
	*accessory.Accessory
	Fan *service.Fan
}

func (acc *AccessoryFanSwitch) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryFanSwitch) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryFanSwitch) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryFanSwitch) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryFanSwitch) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryFanSwitch return AccessoryFanSwitch (args... are not used)
func NewAccessoryFanSwitch(info accessory.Info, args ...interface{}) *AccessoryFanSwitch {
	acc := AccessoryFanSwitch{}
	acc.Accessory = accessory.New(info, accessory.TypeFan)
	acc.Fan = service.NewFan()
	acc.AddService(acc.Fan.Service)
	return &acc
}

func (acc *AccessoryFanSwitch) OnValuesRemoteUpdates(fn func()) {
	acc.Fan.On.OnValueRemoteUpdate(func(bool) { fn() })
}

func (acc *AccessoryFanSwitch) OnValuesRemoteUpdatesPrint() {
	acc.Fan.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%T - %s] remote update on: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
