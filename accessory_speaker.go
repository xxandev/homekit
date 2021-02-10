package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryTypeSpeaker -
const AccessoryTypeSpeaker accessory.AccessoryType = 26

//AccessorySpeaker struct
type AccessorySpeaker struct {
	*accessory.Accessory
	Speaker *service.Speaker
}

//NewAccessorySpeaker returns AccessorySwitch (args... are not used)
func NewAccessorySpeaker(info accessory.Info, args ...interface{}) *AccessorySpeaker {
	acc := AccessorySpeaker{}
	acc.Accessory = accessory.New(info, AccessoryTypeSpeaker)
	acc.Speaker = service.NewSpeaker()
	acc.AddService(acc.Speaker.Service)
	return &acc
}
