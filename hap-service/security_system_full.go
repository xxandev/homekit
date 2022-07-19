package haps

import (
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
)

//SecuritySystemMultifunc
//	◈ SecuritySystemCurrentState
//	◈ SecuritySystemTargetState
//	◇ SecuritySystemAlarmType
//	◇ StatusFault
//	◇ StatusTampered
type SecuritySystemFull struct {
	*service.S
	SecuritySystemCurrentState *characteristic.SecuritySystemCurrentState
	SecuritySystemTargetState  *characteristic.SecuritySystemTargetState
	SecuritySystemAlarmType    *characteristic.SecuritySystemAlarmType
	StatusFault                *characteristic.StatusFault
	StatusTampered             *characteristic.StatusTampered
}

//NewSecuritySystemMultifunc return *SecuritySystemMultifunc
func NewSecuritySystemMultifunc() *SecuritySystemFull {
	svc := SecuritySystemFull{}
	svc.S = service.New(service.TypeSecuritySystem)

	svc.SecuritySystemCurrentState = characteristic.NewSecuritySystemCurrentState()
	svc.AddC(svc.SecuritySystemCurrentState.C)

	svc.SecuritySystemTargetState = characteristic.NewSecuritySystemTargetState()
	svc.AddC(svc.SecuritySystemTargetState.C)

	svc.SecuritySystemAlarmType = characteristic.NewSecuritySystemAlarmType()
	svc.AddC(svc.SecuritySystemAlarmType.C)

	svc.StatusFault = characteristic.NewStatusFault()
	svc.AddC(svc.StatusFault.C)

	svc.StatusTampered = characteristic.NewStatusTampered()
	svc.AddC(svc.StatusTampered.C)

	return &svc
}
