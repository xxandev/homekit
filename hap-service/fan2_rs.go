package haps

import (
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
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
	*service.S
	Active        *characteristic.Active
	RotationSpeed *characteristic.RotationSpeed
}

//NewFanV2Multifunc return *FanV2Multifunc
func NewFanV2Multifunc() *Fan2RS {
	svc := Fan2RS{}
	svc.S = service.New(service.TypeFanV2)

	svc.Active = characteristic.NewActive()
	svc.AddC(svc.Active.C)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddC(svc.RotationSpeed.C)

	return &svc
}
