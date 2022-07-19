package haps

import (
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
)

//HumidifierDehumidifier
//	◈ Active
//	◈ CurrentRelativeHumidity
//	◈ CurrentHumidifierDehumidifierState
//	◈ TargetHumidifierDehumidifierState
//	◇ RelativeHumidityDehumidifierThreshold
//	◇ RelativeHumidityHumidifierThreshold
type HumidifierDehumidifier struct {
	*service.S
	Active                                *characteristic.Active
	CurrentHumidifierDehumidifierState    *characteristic.CurrentHumidifierDehumidifierState
	TargetHumidifierDehumidifierState     *characteristic.TargetHumidifierDehumidifierState
	CurrentRelativeHumidity               *characteristic.CurrentRelativeHumidity
	RelativeHumidityDehumidifierThreshold *characteristic.RelativeHumidityDehumidifierThreshold
	RelativeHumidityHumidifierThreshold   *characteristic.RelativeHumidityHumidifierThreshold
}

//NewHumidifierDehumidifier return *HumidifierDehumidifier
func NewHumidifierDehumidifier() *HumidifierDehumidifier {
	svc := HumidifierDehumidifier{}
	svc.S = service.New(service.TypeHumidifierDehumidifier)

	svc.Active = characteristic.NewActive()
	svc.AddC(svc.Active.C)

	svc.CurrentHumidifierDehumidifierState = characteristic.NewCurrentHumidifierDehumidifierState()
	svc.AddC(svc.CurrentHumidifierDehumidifierState.C)

	svc.TargetHumidifierDehumidifierState = characteristic.NewTargetHumidifierDehumidifierState()
	svc.AddC(svc.TargetHumidifierDehumidifierState.C)

	svc.CurrentRelativeHumidity = characteristic.NewCurrentRelativeHumidity()
	svc.AddC(svc.CurrentRelativeHumidity.C)

	svc.RelativeHumidityDehumidifierThreshold = characteristic.NewRelativeHumidityDehumidifierThreshold()
	svc.AddC(svc.RelativeHumidityDehumidifierThreshold.C)

	svc.RelativeHumidityHumidifierThreshold = characteristic.NewRelativeHumidityHumidifierThreshold()
	svc.AddC(svc.RelativeHumidityHumidifierThreshold.C)

	return &svc
}
