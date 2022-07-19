package homekit

import (
	"fmt"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryFaucet struct
type AccessoryFaucet struct {
	*accessory.A
	// Faucet *service.Faucet
	Valve *service.Valve
}

func (acc *AccessoryFaucet) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryFaucet) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryFaucet) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryFaucet) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryFaucet) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryFaucet) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryFaucet return AccessoryFaucet (args... are not used)
func NewAccessoryFaucet(info accessory.Info, args ...interface{}) *AccessoryFaucet {
	acc := AccessoryFaucet{}
	acc.A = accessory.New(info, accessory.TypeFaucet)
	// acc.Faucet = service.NewFaucet()
	acc.Valve = service.NewValve()

	acc.Valve.ValveType.SetValue(0)

	// acc.AddService(acc.Faucet.Service)
	acc.AddS(acc.Valve.S)
	return &acc
}

func (acc *AccessoryFaucet) OnValuesRemoteUpdates(fn func()) {
	acc.Valve.Active.OnValueRemoteUpdate(func(int) { fn() })
	// acc.Valve.InUse.OnValueRemoteUpdate(func(int) { fn() })
}

func (acc *AccessoryFaucet) OnExample() {
	acc.Valve.Active.OnValueRemoteUpdate(func(v int) {
		acc.Valve.InUse.SetValue(v)
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update active: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
}
