package homekit

import (
	"fmt"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryIrrigation struct
type AccessoryIrrigation struct {
	*accessory.Accessory
	Valve *service.Valve
}

func (acc *AccessoryIrrigation) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryIrrigation) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryIrrigation) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryIrrigation) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryIrrigation) GetAccessory() *accessory.Accessory {
	return acc.Accessory
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

func (acc *AccessoryIrrigation) OnValuesRemoteUpdatesPrint() {
	acc.Valve.Active.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update active: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
