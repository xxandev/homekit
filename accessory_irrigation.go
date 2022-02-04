package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryIrrigation struct
type AccessoryIrrigation struct {
	*accessory.Accessory
	Valve *service.Valve
}

//NewAccessoryIrrigation return AccessoryIrrigation (args... are not used)
func NewAccessoryIrrigation(info accessory.Info, args ...interface{}) *AccessoryIrrigation {
	acc := AccessoryIrrigation{}
	acc.Accessory = accessory.New(info, accessory.TypeSprinklers)
	acc.Valve = service.NewValve()
	acc.Valve.ValveType.SetValue(1)
	acc.AddService(acc.Valve.Service)
	return &acc
}

func (acc *AccessoryIrrigation) OnValuesRemoteUpdates(fn func()) {
	acc.Valve.Active.OnValueRemoteUpdate(func(int) { fn() })
	// acc.Valve.InUse.OnValueRemoteUpdate(func(int) { fn() })
}
