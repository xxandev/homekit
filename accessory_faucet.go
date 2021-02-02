package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryFaucet struct
type AccessoryFaucet struct {
	*accessory.Accessory
	Faucet *service.Faucet
}

//NewAccessoryFaucet return AccessoryFaucet (args... are not used)
func NewAccessoryFaucet(info accessory.Info, args ...interface{}) *AccessoryFaucet {
	acc := AccessoryFaucet{}
	acc.Accessory = accessory.New(info, accessory.TypeFaucets)
	acc.Faucet = service.NewFaucet()
	acc.AddService(acc.Faucet.Service)
	return &acc
}
