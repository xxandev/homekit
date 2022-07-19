package haps

import (
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
)

//FanRS
//	◈ On
//	◇ RotationSpeed
type FanRS struct {
	*service.S
	On            *characteristic.On
	RotationSpeed *characteristic.RotationSpeed
}

//NewFanRS return *FanRS
func NewFanRS() *FanRS {
	svc := FanRS{}
	svc.S = service.New(service.TypeFan)

	svc.On = characteristic.NewOn()
	svc.AddC(svc.On.C)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddC(svc.RotationSpeed.C)

	return &svc
}
