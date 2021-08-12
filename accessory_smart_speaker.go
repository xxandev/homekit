package homekit

import (
	haps "github.com/alpr777/homekit/hap-service"
	"github.com/brutella/hc/accessory"
)

//AccessorySmartSpeaker struct
type AccessorySmartSpeaker struct {
	*accessory.Accessory
	SmartSpeaker *haps.SmartSpeaker
}

//NewAccessorySmartSpeaker returns AccessorySmartSpeaker (args... are not used)
func NewAccessorySmartSpeaker(info accessory.Info, args ...interface{}) *AccessorySmartSpeaker {
	acc := AccessorySmartSpeaker{}
	acc.Accessory = accessory.New(info, AccessoryTypeSpeaker)
	acc.SmartSpeaker = haps.NewSmartSpeaker()
	acc.AddService(acc.SmartSpeaker.Service)
	return &acc
}
