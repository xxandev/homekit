package homekit

import (
	"fmt"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryFanSwitch struct
type AccessoryFan2Switch struct {
	*accessory.A
	Fan2 *service.FanV2
}

func (acc *AccessoryFan2Switch) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryFan2Switch) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryFan2Switch) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryFan2Switch) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryFan2Switch) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryFan2Switch) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryFanSwitch return AccessoryFanSwitch (args... are not used)
func NewAccessoryFan2Switch(info accessory.Info, args ...interface{}) *AccessoryFan2Switch {
	acc := AccessoryFan2Switch{}
	acc.A = accessory.New(info, accessory.TypeFan)
	acc.Fan2 = service.NewFanV2()
	acc.AddS(acc.Fan2.S)
	return &acc
}

func (acc *AccessoryFan2Switch) OnValuesRemoteUpdates(fn func()) {
	acc.Fan2.Active.OnValueRemoteUpdate(func(int) { fn() })
}

func (acc *AccessoryFan2Switch) OnExample() {
	acc.Fan2.Active.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update active: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
}
