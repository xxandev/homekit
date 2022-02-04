package homekit

import (
	"fmt"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryGate struct
type AccessoryGate struct {
	*accessory.Accessory
	GarageDoorOpener *service.GarageDoorOpener
}

func (acc *AccessoryGate) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryGate) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryGate) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryGate) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryGate) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryGate return AccessoryGate (args... are not used)
func NewAccessoryGate(info accessory.Info, args ...interface{}) *AccessoryGate {
	acc := AccessoryGate{}
	acc.Accessory = accessory.New(info, accessory.TypeGarageDoorOpener)
	acc.GarageDoorOpener = service.NewGarageDoorOpener()
	acc.AddService(acc.GarageDoorOpener.Service)
	return &acc
}

func (acc *AccessoryGate) OnValuesRemoteUpdates(fn func()) {
	acc.GarageDoorOpener.TargetDoorState.OnValueRemoteUpdate(func(int) { fn() })
	// acc.GarageDoorOpener.ObstructionDetected.OnValueRemoteUpdate(func(bool) { fn() })
}

func (acc *AccessoryGate) OnValuesRemoteUpdatesPrint() {
	acc.GarageDoorOpener.TargetDoorState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update target state: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
