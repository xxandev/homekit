package haps

import (
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
)

const TypeInputSource = "D9"

type InputSource struct {
	*service.S

	ConfiguredName         *characteristic.ConfiguredName
	InputSourceType        *characteristic.InputSourceType
	IsConfigured           *characteristic.IsConfigured
	CurrentVisibilityState *characteristic.CurrentVisibilityState
	Identifier             *characteristic.Identifier
	InputDeviceType        *characteristic.InputDeviceType
	TargetVisibilityState  *characteristic.TargetVisibilityState
	Name                   *characteristic.Name
}

func NewInputSource() *InputSource {
	svc := InputSource{}
	svc.S = service.New(TypeInputSource)

	svc.ConfiguredName = characteristic.NewConfiguredName()
	svc.AddC(svc.ConfiguredName.C)

	svc.InputSourceType = characteristic.NewInputSourceType()
	svc.AddC(svc.InputSourceType.C)

	svc.IsConfigured = characteristic.NewIsConfigured()
	svc.AddC(svc.IsConfigured.C)

	svc.CurrentVisibilityState = characteristic.NewCurrentVisibilityState()
	svc.AddC(svc.CurrentVisibilityState.C)

	svc.Identifier = characteristic.NewIdentifier()
	svc.AddC(svc.Identifier.C)

	svc.InputDeviceType = characteristic.NewInputDeviceType()
	svc.AddC(svc.InputDeviceType.C)

	svc.TargetVisibilityState = characteristic.NewTargetVisibilityState()
	svc.AddC(svc.TargetVisibilityState.C)

	svc.Name = characteristic.NewName()
	svc.AddC(svc.Name.C)

	return &svc
}
