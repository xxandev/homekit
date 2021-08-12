package haps

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//HumidifierDehumidifier (+Active, +CurrentRelativeHumidity, +CurrentHumidifierDehumidifierState,
//+TargetHumidifierDehumidifierState, RelativeHumidityDehumidifierThreshold, RelativeHumidityHumidifierThreshold)
type HumidifierDehumidifier struct {
	*service.Service

	Active                                *characteristic.Active
	CurrentHumidifierDehumidifierState    *characteristic.CurrentHumidifierDehumidifierState
	TargetHumidifierDehumidifierState     *characteristic.TargetHumidifierDehumidifierState
	CurrentRelativeHumidity               *characteristic.CurrentRelativeHumidity
	RelativeHumidityDehumidifierThreshold *characteristic.RelativeHumidityDehumidifierThreshold
	RelativeHumidityHumidifierThreshold   *characteristic.RelativeHumidityHumidifierThreshold
	// RotationSpeed *characteristic.RotationSpeed
	// SwingMode *characteristic.SwingMode
	// WaterLevel *characteristic.WaterLevel
	// LockPhysicalControls  *characteristic.LockPhysicalControls
}

//NewHumidifierDehumidifier return *HumidifierDehumidifier
func NewHumidifierDehumidifier() *HumidifierDehumidifier {
	svc := HumidifierDehumidifier{}
	svc.Service = service.New(service.TypeHumidifierDehumidifier)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.CurrentHumidifierDehumidifierState = characteristic.NewCurrentHumidifierDehumidifierState()
	svc.AddCharacteristic(svc.CurrentHumidifierDehumidifierState.Characteristic)

	svc.TargetHumidifierDehumidifierState = characteristic.NewTargetHumidifierDehumidifierState()
	svc.AddCharacteristic(svc.TargetHumidifierDehumidifierState.Characteristic)

	svc.CurrentRelativeHumidity = characteristic.NewCurrentRelativeHumidity()
	svc.AddCharacteristic(svc.CurrentRelativeHumidity.Characteristic)

	svc.RelativeHumidityDehumidifierThreshold = characteristic.NewRelativeHumidityDehumidifierThreshold()
	svc.AddCharacteristic(svc.RelativeHumidityDehumidifierThreshold.Characteristic)

	svc.RelativeHumidityHumidifierThreshold = characteristic.NewRelativeHumidityHumidifierThreshold()
	svc.AddCharacteristic(svc.RelativeHumidityHumidifierThreshold.Characteristic)

	return &svc
}
