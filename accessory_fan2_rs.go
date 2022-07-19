package homekit

import (
	"fmt"

	"github.com/brutella/hap/accessory"
	haps "github.com/xxandev/homekit/hap-service"
)

//AccessoryFanV2Multifunc struct
type AccessoryFan2RS struct {
	*accessory.A
	Fan2 *haps.Fan2RS
}

func (acc *AccessoryFan2RS) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryFan2RS) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryFan2RS) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryFan2RS) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryFan2RS) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryFan2RS) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryFanV2Multifunc return AccessoryFanV2Multifunc (args... are not used)
func NewAccessoryFan2RS(info accessory.Info, args ...interface{}) *AccessoryFan2RS {
	acc := AccessoryFan2RS{}
	acc.A = accessory.New(info, accessory.TypeFan)
	acc.Fan2 = haps.NewFanV2Multifunc()
	acc.AddS(acc.Fan2.S)
	return &acc
}

func (acc *AccessoryFan2RS) OnValuesRemoteUpdates(fn func()) {
	acc.Fan2.Active.OnValueRemoteUpdate(func(int) { fn() })
	acc.Fan2.RotationSpeed.OnValueRemoteUpdate(func(float64) { fn() })
}

func (acc *AccessoryFan2RS) OnExample() {
	acc.Fan2.Active.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update active: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
	acc.Fan2.RotationSpeed.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update rotation speed: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
}
