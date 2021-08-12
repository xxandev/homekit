package haps

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//TypeSmartSpeaker - 00000228-0000-1000-8000-0026BB765291
const TypeSmartSpeaker string = "228"

//SmartSpeaker (+CurrentMediaState, +TargetMediaState, Name, ConfiguredName, Mute, Volume)
type SmartSpeaker struct {
	*service.Service
	CurrentMediaState *characteristic.CurrentMediaState
	TargetMediaState  *characteristic.TargetMediaState
	Name              *characteristic.Name
	ConfiguredName    *characteristic.ConfiguredName
	Mute              *characteristic.Mute
	Volume            *characteristic.Volume
}

//NewSmartSpeaker return *SmartSpeaker
func NewSmartSpeaker() *SmartSpeaker {
	svc := SmartSpeaker{}
	svc.Service = service.New(TypeSmartSpeaker)

	svc.TargetMediaState = characteristic.NewTargetMediaState()
	svc.AddCharacteristic(svc.TargetMediaState.Characteristic)

	svc.CurrentMediaState = characteristic.NewCurrentMediaState()
	svc.AddCharacteristic(svc.CurrentMediaState.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	svc.ConfiguredName = characteristic.NewConfiguredName()
	svc.AddCharacteristic(svc.ConfiguredName.Characteristic)

	svc.Mute = characteristic.NewMute()
	svc.AddCharacteristic(svc.Mute.Characteristic)

	svc.Volume = characteristic.NewVolume()
	svc.AddCharacteristic(svc.Volume.Characteristic)

	return &svc
}
