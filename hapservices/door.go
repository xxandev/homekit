package hapservices

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//Door -
type Door struct {
	*service.Service
	CurrentPosition     *characteristic.CurrentPosition
	PositionState       *characteristic.PositionState
	TargetPosition      *characteristic.TargetPosition
	ObstructionDetected *characteristic.ObstructionDetected
}

//NewDoor -
func NewDoor() *Door {
	svc := Door{}
	svc.Service = service.New(service.TypeDoor)

	svc.CurrentPosition = characteristic.NewCurrentPosition()
	svc.AddCharacteristic(svc.CurrentPosition.Characteristic)

	svc.PositionState = characteristic.NewPositionState()
	svc.AddCharacteristic(svc.PositionState.Characteristic)

	svc.TargetPosition = characteristic.NewTargetPosition()
	svc.AddCharacteristic(svc.TargetPosition.Characteristic)

	svc.ObstructionDetected = characteristic.NewObstructionDetected()
	svc.AddCharacteristic(svc.ObstructionDetected.Characteristic)

	return &svc
}
