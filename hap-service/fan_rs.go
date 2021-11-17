package haps

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//FanRS
//	◈ On
//	◇ RotationSpeed
type FanRS struct {
	*service.Service
	On            *characteristic.On
	RotationSpeed *characteristic.RotationSpeed
}

//NewFanRS return *FanRS
func NewFanRS() *FanRS {
	svc := FanRS{}
	svc.Service = service.New(service.TypeFan)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddCharacteristic(svc.RotationSpeed.Characteristic)

	return &svc
}
