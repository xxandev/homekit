package haps

import (
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
)

//Door
//	◈ CurrentPosition
//	◈ TargetPosition
//	◈ PositionState
//	◇ ObstructionDetected
type Door struct {
	*service.S
	CurrentPosition     *characteristic.CurrentPosition
	TargetPosition      *characteristic.TargetPosition
	PositionState       *characteristic.PositionState
	ObstructionDetected *characteristic.ObstructionDetected
}

//NewDoor return *Door
func NewDoor() *Door {
	svc := Door{}
	svc.S = service.New(service.TypeDoor)

	svc.CurrentPosition = characteristic.NewCurrentPosition()
	svc.AddC(svc.CurrentPosition.C)

	svc.PositionState = characteristic.NewPositionState()
	svc.AddC(svc.PositionState.C)

	svc.TargetPosition = characteristic.NewTargetPosition()
	svc.AddC(svc.TargetPosition.C)

	svc.ObstructionDetected = characteristic.NewObstructionDetected()
	svc.AddC(svc.ObstructionDetected.C)

	return &svc
}
