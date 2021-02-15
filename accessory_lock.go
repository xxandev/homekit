package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryLock struct
type AccessoryLock struct {
	*accessory.Accessory
	LockMechanism *service.LockMechanism
}

//NewAccessoryLock return AccessoryDoorLock (args... are not used)
func NewAccessoryLock(info accessory.Info, args ...interface{}) *AccessoryLock {
	acc := AccessoryLock{}
	acc.Accessory = accessory.New(info, accessory.TypeDoorLock)
	acc.LockMechanism = service.NewLockMechanism()
	acc.AddService(acc.LockMechanism.Service)
	return &acc
}
