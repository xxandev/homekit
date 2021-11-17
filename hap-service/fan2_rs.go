package haps

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//Fan2RS
//	◈ Active
//	◇ CurrentFanState
//	◇ TargetFanState
//	◇ RotationDirection
//	◇ RotationSpeed
//	◇ SwingMode
//	◇ LockPhysicalControls
type Fan2RS struct {
	*service.Service
	Active        *characteristic.Active
	RotationSpeed *characteristic.RotationSpeed
}

//NewFanV2Multifunc return *FanV2Multifunc
func NewFanV2Multifunc() *Fan2RS {
	svc := Fan2RS{}
	svc.Service = service.New(service.TypeFanV2)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddCharacteristic(svc.RotationSpeed.Characteristic)

	return &svc
}
