package hapservices

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//AirPurifier (+Active, +CurrentAirPurifierState, +TargetAirPurifierState, RotationSpeed)
type AirPurifier struct {
	*service.Service
	Active                  *characteristic.Active
	CurrentAirPurifierState *characteristic.CurrentAirPurifierState
	TargetAirPurifierState  *characteristic.TargetAirPurifierState
	RotationSpeed           *characteristic.RotationSpeed
}

//NewAirPurifier return *AirPurifier
func NewAirPurifier() *AirPurifier {
	svc := AirPurifier{}
	svc.Service = service.New(service.TypeAirPurifier)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.CurrentAirPurifierState = characteristic.NewCurrentAirPurifierState()
	svc.AddCharacteristic(svc.CurrentAirPurifierState.Characteristic)

	svc.TargetAirPurifierState = characteristic.NewTargetAirPurifierState()
	svc.AddCharacteristic(svc.TargetAirPurifierState.Characteristic)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddCharacteristic(svc.RotationSpeed.Characteristic)

	return &svc
}
