package homekit

import (
	"fmt"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryLock struct
type AccessoryLock struct {
	*accessory.Accessory
	LockMechanism *service.LockMechanism
}

func (acc *AccessoryLock) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryLock) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryLock) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryLock) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryLock) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryLock return AccessoryDoorLock (args... are not used)
func NewAccessoryLock(info accessory.Info, args ...interface{}) *AccessoryLock {
	acc := AccessoryLock{}
	acc.Accessory = accessory.New(info, accessory.TypeDoorLock)
	acc.LockMechanism = service.NewLockMechanism()
	acc.AddService(acc.LockMechanism.Service)
	return &acc
}

func (acc *AccessoryLock) OnValuesRemoteUpdates(fn func()) {
	acc.LockMechanism.LockTargetState.OnValueRemoteUpdate(func(int) { fn() })
}

func (acc *AccessoryLock) OnExample() {
	acc.LockMechanism.LockTargetState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update target state: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
