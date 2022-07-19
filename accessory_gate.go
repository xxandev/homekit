package homekit

import (
	"fmt"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryGate struct
type AccessoryGate struct {
	*accessory.A
	GarageDoorOpener *service.GarageDoorOpener
}

func (acc *AccessoryGate) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryGate) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryGate) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryGate) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryGate) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryGate) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryGate return AccessoryGate (args... are not used)
func NewAccessoryGate(info accessory.Info, args ...interface{}) *AccessoryGate {
	acc := AccessoryGate{}
	acc.A = accessory.New(info, accessory.TypeGarageDoorOpener)
	acc.GarageDoorOpener = service.NewGarageDoorOpener()
	acc.AddS(acc.GarageDoorOpener.S)
	return &acc
}

func (acc *AccessoryGate) OnValuesRemoteUpdates(fn func()) {
	acc.GarageDoorOpener.TargetDoorState.OnValueRemoteUpdate(func(int) { fn() })
	// acc.GarageDoorOpener.ObstructionDetected.OnValueRemoteUpdate(func(bool) { fn() })
}

func (acc *AccessoryGate) OnExample() {
	acc.GarageDoorOpener.TargetDoorState.OnValueRemoteUpdate(func(v int) {
		acc.GarageDoorOpener.CurrentDoorState.SetValue(v)
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update target state: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
}
