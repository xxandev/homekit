package homekit

import (
	"github.com/alpr777/homekit/hapservices"
	"github.com/brutella/hc/accessory"
)

//AccessorySmartSpeaker struct
type AccessorySmartSpeaker struct {
	*accessory.Accessory
	SmartSpeaker *hapservices.SmartSpeaker
}

//NewAccessorySmartSpeaker returns AccessorySmartSpeaker (args... are not used)
func NewAccessorySmartSpeaker(info accessory.Info, args ...interface{}) *AccessorySmartSpeaker {
	acc := AccessorySmartSpeaker{}
	acc.Accessory = accessory.New(info, AccessoryTypeSpeaker)
	acc.SmartSpeaker = hapservices.NewSmartSpeaker()
	acc.AddService(acc.SmartSpeaker.Service)
	return &acc
}
