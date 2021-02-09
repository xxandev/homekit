package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryDoorLock struct
type AccessoryDoorLock struct {
	*accessory.Accessory
	LockMechanism *service.LockMechanism
}

//NewAccessoryDoorLock return AccessoryDoorLock (args... are not used)
func NewAccessoryDoorLock(info accessory.Info, args ...interface{}) *AccessoryDoorLock {
	acc := AccessoryDoorLock{}
	acc.Accessory = accessory.New(info, accessory.TypeDoorLock)
	acc.LockMechanism = service.NewLockMechanism()
	acc.AddService(acc.LockMechanism.Service)
	return &acc
}
