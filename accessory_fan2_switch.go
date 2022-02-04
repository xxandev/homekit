package homekit

import (
	"fmt"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryFanSwitch struct
type AccessoryFan2Switch struct {
	*accessory.Accessory
	Fan2 *service.FanV2
}

func (acc *AccessoryFan2Switch) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryFan2Switch) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryFan2Switch) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryFan2Switch) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryFan2Switch) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryFanSwitch return AccessoryFanSwitch (args... are not used)
func NewAccessoryFan2Switch(info accessory.Info, args ...interface{}) *AccessoryFan2Switch {
	acc := AccessoryFan2Switch{}
	acc.Accessory = accessory.New(info, accessory.TypeFan)
	acc.Fan2 = service.NewFanV2()
	acc.AddService(acc.Fan2.Service)
	return &acc
}

func (acc *AccessoryFan2Switch) OnValuesRemoteUpdates(fn func()) {
	acc.Fan2.Active.OnValueRemoteUpdate(func(int) { fn() })
}

func (acc *AccessoryFan2Switch) OnValuesRemoteUpdatesPrint() {
	acc.Fan2.Active.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update active: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
