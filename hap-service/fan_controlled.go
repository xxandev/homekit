package haps

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//FanControlled (+On, RotationSpeed)
type FanControlled struct {
	*service.Service
	On            *characteristic.On
	RotationSpeed *characteristic.RotationSpeed
}

//NewFanControlled return *FanControlled
func NewFanControlled() *FanControlled {
	svc := FanControlled{}
	svc.Service = service.New(service.TypeFan)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddCharacteristic(svc.RotationSpeed.Characteristic)

	return &svc
}
