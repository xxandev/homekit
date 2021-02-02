package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySwitch struct
type AccessorySwitch struct {
	*accessory.Accessory
	Switch *service.Switch
}

// NewAccessorySwitch returns AccessorySwitch (args... are not used)
func NewAccessorySwitch(info accessory.Info, args ...interface{}) *AccessorySwitch {
	acc := AccessorySwitch{}
	acc.Accessory = accessory.New(info, accessory.TypeSwitch)
	acc.Switch = service.NewSwitch()
	acc.AddService(acc.Switch.Service)
	return &acc
}
