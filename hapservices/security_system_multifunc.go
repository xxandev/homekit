package hapservices

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//SecuritySystemMultifunc -
type SecuritySystemMultifunc struct {
	*service.Service
	SecuritySystemCurrentState *characteristic.SecuritySystemCurrentState
	SecuritySystemTargetState  *characteristic.SecuritySystemTargetState
	SecuritySystemAlarmType    *characteristic.SecuritySystemAlarmType
	StatusFault                *characteristic.StatusFault
	StatusTampered             *characteristic.StatusTampered
}

//NewSecuritySystemMultifunc -
func NewSecuritySystemMultifunc() *SecuritySystemMultifunc {
	svc := SecuritySystemMultifunc{}
	svc.Service = service.New(service.TypeSecuritySystem)

	svc.SecuritySystemCurrentState = characteristic.NewSecuritySystemCurrentState()
	svc.AddCharacteristic(svc.SecuritySystemCurrentState.Characteristic)

	svc.SecuritySystemTargetState = characteristic.NewSecuritySystemTargetState()
	svc.AddCharacteristic(svc.SecuritySystemTargetState.Characteristic)

	svc.SecuritySystemAlarmType = characteristic.NewSecuritySystemAlarmType()
	svc.AddCharacteristic(svc.SecuritySystemAlarmType.Characteristic)

	svc.StatusFault = characteristic.NewStatusFault()
	svc.AddCharacteristic(svc.StatusFault.Characteristic)

	svc.StatusTampered = characteristic.NewStatusTampered()
	svc.AddCharacteristic(svc.StatusTampered.Characteristic)

	return &svc
}
