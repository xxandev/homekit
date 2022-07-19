package haps

import (
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
)

//TelevisionSpeaker
//	◈ Mute
//	◇ Active
//	◇ Volume
//	◇ VolumeControlType
//	◇ VolumeSelector
type TelevisionSpeaker struct {
	*service.S

	Mute *characteristic.Mute
	//Active            *characteristic.Active
	Volume            *characteristic.Volume
	VolumeControlType *characteristic.VolumeControlType
	VolumeSelector    *characteristic.VolumeSelector
}

//NewTelevisionSpeaker return *TelevisionSpeaker
func NewTelevisionSpeaker() *TelevisionSpeaker {
	svc := TelevisionSpeaker{}
	svc.S = service.New(service.TypeSpeaker)

	svc.Mute = characteristic.NewMute()
	svc.AddC(svc.Mute.C)

	// svc.Active = characteristic.NewActive()
	// svc.AddCharacteristic(svc.Active.Characteristic)

	svc.Volume = characteristic.NewVolume()
	svc.AddC(svc.Volume.C)

	svc.VolumeControlType = characteristic.NewVolumeControlType()
	svc.AddC(svc.VolumeControlType.C)

	svc.VolumeSelector = characteristic.NewVolumeSelector()
	svc.AddC(svc.VolumeSelector.C)

	return &svc
}
