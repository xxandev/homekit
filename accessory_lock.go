package homekit

import (
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

//NewAccessoryLock return *DoorLock.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessoryLock(info accessory.Info, args ...interface{}) *AccessoryLock {
	acc := AccessoryLock{}
	acc.A = accessory.New(info, accessory.TypeDoorLock)
	acc.LockMechanism = service.NewLockMechanism()
	acc.AddS(acc.LockMechanism.S)
	return &acc
}

//NewAccLock return *DoorLock.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccLock(id uint64, info accessory.Info, args ...interface{}) *AccessoryLock {
	acc := AccessoryLock{}
	acc.A = accessory.New(info, accessory.TypeDoorLock)
	acc.LockMechanism = service.NewLockMechanism()
	acc.AddS(acc.LockMechanism.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessoryLock) OnValuesRemoteUpdates(fn func()) {
	acc.LockMechanism.LockTargetState.OnValueRemoteUpdate(func(int) { fn() })
}
