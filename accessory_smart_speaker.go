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

func (acc *AccessorySmartSpeaker) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessorySmartSpeaker) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessorySmartSpeaker) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessorySmartSpeaker) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessorySmartSpeaker) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessorySmartSpeaker returns AccessorySmartSpeaker (args... are not used)
func NewAccessorySmartSpeaker(info accessory.Info, args ...interface{}) *AccessorySmartSpeaker {
	acc := AccessorySmartSpeaker{}
	acc.Accessory = accessory.New(info, AccessoryTypeSpeaker)
	acc.SmartSpeaker = haps.NewSmartSpeaker()
	acc.AddService(acc.SmartSpeaker.Service)
	return &acc
}

func (acc *AccessorySmartSpeaker) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySmartSpeaker) OnValuesRemoteUpdatesPrint()     {}
