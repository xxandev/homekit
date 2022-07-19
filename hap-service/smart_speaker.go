package haps

import (
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
)

//TypeSmartSpeaker - 00000228-0000-1000-8000-0026BB765291
const TypeSmartSpeaker string = "228"

//SmartSpeaker
//	◈ CurrentMediaState
//	◈ TargetMediaState
//	◇ Name
//	◇ ConfiguredName
//	◇ Mute
//	◇ Volume
type SmartSpeaker struct {
	*service.S
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
	svc.S = service.New(TypeSmartSpeaker)

	svc.TargetMediaState = characteristic.NewTargetMediaState()
	svc.AddC(svc.TargetMediaState.C)

	svc.CurrentMediaState = characteristic.NewCurrentMediaState()
	svc.AddC(svc.CurrentMediaState.C)

	svc.Name = characteristic.NewName()
	svc.AddC(svc.Name.C)

	svc.ConfiguredName = characteristic.NewConfiguredName()
	svc.AddC(svc.ConfiguredName.C)

	svc.Mute = characteristic.NewMute()
	svc.AddC(svc.Mute.C)

	svc.Volume = characteristic.NewVolume()
	svc.AddC(svc.Volume.C)

	return &svc
}
