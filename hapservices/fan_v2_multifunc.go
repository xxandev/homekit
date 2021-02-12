package hapservices

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//FanV2Multifunc -
type FanV2Multifunc struct {
	*service.Service
	Active               *characteristic.Active
	CurrentFanState      *characteristic.CurrentFanState
	TargetFanState       *characteristic.TargetFanState
	RotationDirection    *characteristic.RotationDirection
	RotationSpeed        *characteristic.RotationSpeed
	SwingMode            *characteristic.SwingMode
	LockPhysicalControls *characteristic.LockPhysicalControls
}

//NewFanV2Multifunc -
func NewFanV2Multifunc() *FanV2Multifunc {
	svc := FanV2Multifunc{}
	svc.Service = service.New(service.TypeFanV2)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.CurrentFanState = characteristic.NewCurrentFanState()
	svc.AddCharacteristic(svc.CurrentFanState.Characteristic)

	svc.TargetFanState = characteristic.NewTargetFanState()
	svc.AddCharacteristic(svc.TargetFanState.Characteristic)

	svc.RotationDirection = characteristic.NewRotationDirection()
	svc.AddCharacteristic(svc.RotationDirection.Characteristic)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddCharacteristic(svc.RotationSpeed.Characteristic)

	svc.SwingMode = characteristic.NewSwingMode()
	svc.AddCharacteristic(svc.SwingMode.Characteristic)

	svc.LockPhysicalControls = characteristic.NewLockPhysicalControls()
	svc.AddCharacteristic(svc.LockPhysicalControls.Characteristic)

	return &svc
}
