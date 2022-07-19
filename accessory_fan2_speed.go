package homekit

import (
	"fmt"

	"github.com/brutella/hc/accessory"
	haps "github.com/xxandev/homekit/hap-service"
)

//AccessoryFanV2Multifunc struct
type AccessoryFan2Speed struct {
	*accessory.Accessory
	Fan2 *haps.Fan2RS
}

func (acc *AccessoryFan2Speed) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryFan2Speed) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryFan2Speed) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryFan2Speed) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryFan2Speed) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryFanV2Multifunc return AccessoryFanV2Multifunc (args... are not used)
func NewAccessoryFan2Speed(info accessory.Info, args ...interface{}) *AccessoryFan2Speed {
	acc := AccessoryFan2Speed{}
	acc.Accessory = accessory.New(info, accessory.TypeFan)
	acc.Fan2 = haps.NewFanV2Multifunc()
	acc.AddService(acc.Fan2.Service)
	return &acc
}

func (acc *AccessoryFan2Speed) OnValuesRemoteUpdates(fn func()) {
	acc.Fan2.Active.OnValueRemoteUpdate(func(int) { fn() })
	acc.Fan2.RotationSpeed.OnValueRemoteUpdate(func(float64) { fn() })
}

func (acc *AccessoryFan2Speed) OnExample() {
	acc.Fan2.Active.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update active: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.Fan2.RotationSpeed.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%T - %s] remote update rotation speed: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
