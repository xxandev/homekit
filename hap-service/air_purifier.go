package haps

import (
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
)

//AirPurifier
//	◈ Active
//	◈ CurrentAirPurifierState
//	◈ TargetAirPurifierState
//	◇ RotationSpeed
type AirPurifier struct {
	*service.S
	Active                  *characteristic.Active
	CurrentAirPurifierState *characteristic.CurrentAirPurifierState
	TargetAirPurifierState  *characteristic.TargetAirPurifierState
	RotationSpeed           *characteristic.RotationSpeed
}

//NewAirPurifier return *AirPurifier
func NewAirPurifier() *AirPurifier {
	svc := AirPurifier{}
	svc.S = service.New(service.TypeAirPurifier)

	svc.Active = characteristic.NewActive()
	svc.AddC(svc.Active.C)

	svc.CurrentAirPurifierState = characteristic.NewCurrentAirPurifierState()
	svc.AddC(svc.CurrentAirPurifierState.C)

	svc.TargetAirPurifierState = characteristic.NewTargetAirPurifierState()
	svc.AddC(svc.TargetAirPurifierState.C)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddC(svc.RotationSpeed.C)

	return &svc
}
