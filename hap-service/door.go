package haps

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//Door (+CurrentPosition, +TargetPosition, +PositionState, ObstructionDetected)
type Door struct {
	*service.Service
	CurrentPosition     *characteristic.CurrentPosition
	TargetPosition      *characteristic.TargetPosition
	PositionState       *characteristic.PositionState
	ObstructionDetected *characteristic.ObstructionDetected
}

//NewDoor return *Door
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
