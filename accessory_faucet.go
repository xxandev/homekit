package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryFaucet struct
type AccessoryFaucet struct {
	*accessory.Accessory
	// Faucet *service.Faucet
	Valve *service.Valve
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
