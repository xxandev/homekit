package hapservices

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//TelevisionSpeaker -
type TelevisionSpeaker struct {
	*service.Service

	Mute              *characteristic.Mute
	Active            *characteristic.Active
	Volume            *characteristic.Volume
	VolumeControlType *characteristic.VolumeControlType
	VolumeSelector    *characteristic.VolumeSelector
}

//NewTelevisionSpeaker -
func NewTelevisionSpeaker() *TelevisionSpeaker {
	svc := TelevisionSpeaker{}
	svc.Service = service.New(service.TypeSpeaker)

	svc.Mute = characteristic.NewMute()
	svc.AddCharacteristic(svc.Mute.Characteristic)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.Volume = characteristic.NewVolume()
	svc.AddCharacteristic(svc.Volume.Characteristic)

	svc.VolumeControlType = characteristic.NewVolumeControlType()
	svc.AddCharacteristic(svc.VolumeControlType.Characteristic)

	svc.VolumeSelector = characteristic.NewVolumeSelector()
	svc.AddCharacteristic(svc.VolumeSelector.Characteristic)

	return &svc
}
