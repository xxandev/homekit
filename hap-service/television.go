package haps

import (
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
)

const TypeTelevision = "D8"

type Television struct {
	*service.S

	Active             *characteristic.Active
	ActiveIdentifier   *characteristic.ActiveIdentifier
	ConfiguredName     *characteristic.ConfiguredName
	SleepDiscoveryMode *characteristic.SleepDiscoveryMode
	Brightness         *characteristic.Brightness
	ClosedCaptions     *characteristic.ClosedCaptions
	DisplayOrder       *characteristic.DisplayOrder
	CurrentMediaState  *characteristic.CurrentMediaState
	TargetMediaState   *characteristic.TargetMediaState
	PictureMode        *characteristic.PictureMode
	PowerModeSelection *characteristic.PowerModeSelection
	RemoteKey          *characteristic.RemoteKey
}

func NewTelevision() *Television {
	svc := Television{}
	svc.S = service.New(TypeTelevision)

	svc.Active = characteristic.NewActive()
	svc.AddC(svc.Active.C)

	svc.ActiveIdentifier = characteristic.NewActiveIdentifier()
	svc.AddC(svc.ActiveIdentifier.C)

	svc.ConfiguredName = characteristic.NewConfiguredName()
	svc.AddC(svc.ConfiguredName.C)

	svc.SleepDiscoveryMode = characteristic.NewSleepDiscoveryMode()
	svc.AddC(svc.SleepDiscoveryMode.C)

	svc.Brightness = characteristic.NewBrightness()
	svc.AddC(svc.Brightness.C)

	svc.ClosedCaptions = characteristic.NewClosedCaptions()
	svc.AddC(svc.ClosedCaptions.C)

	svc.DisplayOrder = characteristic.NewDisplayOrder()
	svc.AddC(svc.DisplayOrder.C)

	svc.CurrentMediaState = characteristic.NewCurrentMediaState()
	svc.AddC(svc.CurrentMediaState.C)

	svc.TargetMediaState = characteristic.NewTargetMediaState()
	svc.AddC(svc.TargetMediaState.C)

	svc.PictureMode = characteristic.NewPictureMode()
	svc.AddC(svc.PictureMode.C)

	svc.PowerModeSelection = characteristic.NewPowerModeSelection()
	svc.AddC(svc.PowerModeSelection.C)

	svc.RemoteKey = characteristic.NewRemoteKey()
	svc.AddC(svc.RemoteKey.C)

	return &svc
}
