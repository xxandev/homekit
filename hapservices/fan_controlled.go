package hapservices

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//FanControlled -
type FanControlled struct {
	*service.Service
	On            *characteristic.On
	RotationSpeed *characteristic.RotationSpeed
}

//NewFanControlled -
func NewFanControlled() *FanControlled {
	svc := FanControlled{}
	svc.Service = service.New(service.TypeFan)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddCharacteristic(svc.RotationSpeed.Characteristic)

	return &svc
}
