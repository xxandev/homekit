package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryLockMechanism struct
type AccessoryLockMechanism struct {
	*accessory.Accessory
	LockMechanism *service.LockMechanism
}

//NewAccessoryLockMechanism return AccessoryLockMechanism (args... are not used)
func NewAccessoryLockMechanism(info accessory.Info, args ...interface{}) *AccessoryLockMechanism {
	acc := AccessoryLockMechanism{}
	acc.Accessory = accessory.New(info, accessory.TypeDoorLock)
	acc.LockMechanism = service.NewLockMechanism()
	acc.AddService(acc.LockMechanism.Service)
	return &acc
}
