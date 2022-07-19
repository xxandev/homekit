package homekit

import (
	"fmt"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryLock struct
type AccessoryLock struct {
	*accessory.A
	LockMechanism *service.LockMechanism
}

func (acc *AccessoryLock) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryLock) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryLock) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryLock) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryLock) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryLock) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryLock return AccessoryDoorLock (args... are not used)
func NewAccessoryLock(info accessory.Info, args ...interface{}) *AccessoryLock {
	acc := AccessoryLock{}
	acc.A = accessory.New(info, accessory.TypeDoorLock)
	acc.LockMechanism = service.NewLockMechanism()
	acc.AddS(acc.LockMechanism.S)
	return &acc
}

func (acc *AccessoryLock) OnValuesRemoteUpdates(fn func()) {
	acc.LockMechanism.LockTargetState.OnValueRemoteUpdate(func(int) { fn() })
}

func (acc *AccessoryLock) OnExample() {
	acc.LockMechanism.LockTargetState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update target position: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
		acc.LockMechanism.LockCurrentState.SetValue(v)
		fmt.Printf("[%[1]T - %[2]v - %[3]v] update current position: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.LockMechanism.LockCurrentState.Value())
	})
}
