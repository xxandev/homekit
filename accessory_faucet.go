package homekit

import (
	"fmt"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryFaucet struct
type AccessoryFaucet struct {
	*accessory.Accessory
	// Faucet *service.Faucet
	Valve *service.Valve
}

func (acc *AccessoryFaucet) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryFaucet) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryFaucet) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryFaucet) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryFaucet) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryFaucet return AccessoryFaucet (args... are not used)
func NewAccessoryFaucet(info accessory.Info, args ...interface{}) *AccessoryFaucet {
	acc := AccessoryFaucet{}
	acc.Accessory = accessory.New(info, accessory.TypeFaucets)
	// acc.Faucet = service.NewFaucet()
	acc.Valve = service.NewValve()

	acc.Valve.ValveType.SetValue(0)

	// acc.AddService(acc.Faucet.Service)
	acc.AddService(acc.Valve.Service)
	return &acc
}

func (acc *AccessoryFaucet) OnValuesRemoteUpdates(fn func()) {
	acc.Valve.Active.OnValueRemoteUpdate(func(int) { fn() })
	// acc.Valve.InUse.OnValueRemoteUpdate(func(int) { fn() })
}

func (acc *AccessoryFaucet) OnExample() {
	acc.Valve.Active.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update active: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
