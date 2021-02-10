package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//AccessoryTypeSpeaker -
const AccessoryTypeSpeaker accessory.AccessoryType = 26

//ServiceSpeaker -
type ServiceSpeaker struct {
	*service.Service
	ConfiguredName    *characteristic.ConfiguredName
	TargetMediaState  *characteristic.TargetMediaState
	CurrentMediaState *characteristic.CurrentMediaState
	Mute              *characteristic.Mute
	Volume            *characteristic.Volume
}

//NewServiceSpeaker -
func NewServiceSpeaker() *ServiceSpeaker {
	svc := ServiceSpeaker{}
	svc.Service = service.New(service.TypeSpeaker)

	svc.ConfiguredName = characteristic.NewConfiguredName()
	svc.AddCharacteristic(svc.ConfiguredName.Characteristic)

	svc.TargetMediaState = characteristic.NewTargetMediaState()
	svc.AddCharacteristic(svc.TargetMediaState.Characteristic)

	svc.CurrentMediaState = characteristic.NewCurrentMediaState()
	svc.AddCharacteristic(svc.CurrentMediaState.Characteristic)

	svc.Mute = characteristic.NewMute()
	svc.AddCharacteristic(svc.Mute.Characteristic)

	svc.Volume = characteristic.NewVolume()
	svc.AddCharacteristic(svc.Volume.Characteristic)

	return &svc
}

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
