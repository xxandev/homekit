package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//ServiceFanV2Multifunc -
type ServiceFanV2Multifunc struct {
	*service.Service
	Active               *characteristic.Active
	CurrentFanState      *characteristic.CurrentFanState
	TargetFanState       *characteristic.TargetFanState
	RotationDirection    *characteristic.RotationDirection
	RotationSpeed        *characteristic.RotationSpeed
	SwingMode            *characteristic.SwingMode
	LockPhysicalControls *characteristic.LockPhysicalControls
}

//NewServiceFanV2Multifunc -
func NewServiceFanV2Multifunc() *ServiceFanV2Multifunc {
	svc := ServiceFanV2Multifunc{}
	svc.Service = service.New(service.TypeFanV2)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.CurrentFanState = characteristic.NewCurrentFanState()
	svc.AddCharacteristic(svc.CurrentFanState.Characteristic)

	svc.TargetFanState = characteristic.NewTargetFanState()
	svc.AddCharacteristic(svc.TargetFanState.Characteristic)

	svc.RotationDirection = characteristic.NewRotationDirection()
	svc.AddCharacteristic(svc.RotationDirection.Characteristic)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddCharacteristic(svc.RotationSpeed.Characteristic)

	svc.SwingMode = characteristic.NewSwingMode()
	svc.AddCharacteristic(svc.SwingMode.Characteristic)

	svc.LockPhysicalControls = characteristic.NewLockPhysicalControls()
	svc.AddCharacteristic(svc.LockPhysicalControls.Characteristic)

	return &svc
}

//AccessoryFanV2Multifunc struct
type AccessoryFanV2Multifunc struct {
	*accessory.Accessory
	FanV2 *ServiceFanV2Multifunc
}

//NewAccessoryFanV2Multifunc return AccessoryFanV2Multifunc (args... are not used)
func NewAccessoryFanV2Multifunc(info accessory.Info, args ...interface{}) *AccessoryFanV2Multifunc {
	acc := AccessoryFanV2Multifunc{}
	acc.Accessory = accessory.New(info, accessory.TypeFan)
	acc.FanV2 = NewServiceFanV2Multifunc()
	acc.AddService(acc.FanV2.Service)
	return &acc
}
