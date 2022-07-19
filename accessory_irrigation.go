package homekit

import (
	"fmt"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryIrrigation struct
type AccessoryIrrigation struct {
	*accessory.A
	Valve *service.Valve
}

func (acc *AccessoryIrrigation) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryIrrigation) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryIrrigation) SetID(id uint64) {
	acc.A.Id = id
}
func (acc *AccessoryIrrigation) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryIrrigation) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryIrrigation) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryIrrigation return AccessoryIrrigation (args... are not used)
func NewAccessoryIrrigation(info accessory.Info, args ...interface{}) *AccessoryIrrigation {
	acc := AccessoryIrrigation{}
	acc.A = accessory.New(info, accessory.TypeSprinkler)
	acc.Valve = service.NewValve()
	acc.Valve.ValveType.SetValue(1)
	acc.AddS(acc.Valve.S)
	return &acc
}

func (acc *AccessoryIrrigation) OnValuesRemoteUpdates(fn func()) {
	acc.Valve.Active.OnValueRemoteUpdate(func(int) { fn() })
	// acc.Valve.InUse.OnValueRemoteUpdate(func(int) { fn() })
}

func (acc *AccessoryIrrigation) OnExample() {
	acc.Valve.Active.OnValueRemoteUpdate(func(v int) {
		acc.Valve.InUse.SetValue(v)
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update active: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
}
