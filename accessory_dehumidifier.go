package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//ServiceHumidifierDehumidifier -
type ServiceHumidifierDehumidifier struct {
	*service.Service

	Active                                *characteristic.Active
	CurrentHumidifierDehumidifierState    *characteristic.CurrentHumidifierDehumidifierState
	TargetHumidifierDehumidifierState     *characteristic.TargetHumidifierDehumidifierState
	CurrentRelativeHumidity               *characteristic.CurrentRelativeHumidity
	TargetRelativeHumidity                *characteristic.TargetRelativeHumidity
	RelativeHumidityDehumidifierThreshold *characteristic.RelativeHumidityDehumidifierThreshold
	RelativeHumidityHumidifierThreshold   *characteristic.RelativeHumidityHumidifierThreshold
}

//NewServiceHumidifierDehumidifier -
func NewServiceHumidifierDehumidifier() *ServiceHumidifierDehumidifier {
	svc := ServiceHumidifierDehumidifier{}
	svc.Service = service.New(service.TypeHumidifierDehumidifier)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.CurrentHumidifierDehumidifierState = characteristic.NewCurrentHumidifierDehumidifierState()
	svc.AddCharacteristic(svc.CurrentHumidifierDehumidifierState.Characteristic)

	svc.TargetHumidifierDehumidifierState = characteristic.NewTargetHumidifierDehumidifierState()
	svc.AddCharacteristic(svc.TargetHumidifierDehumidifierState.Characteristic)

	svc.CurrentRelativeHumidity = characteristic.NewCurrentRelativeHumidity()
	svc.AddCharacteristic(svc.CurrentRelativeHumidity.Characteristic)

	svc.TargetRelativeHumidity = characteristic.NewTargetRelativeHumidity()
	svc.AddCharacteristic(svc.TargetRelativeHumidity.Characteristic)

	svc.RelativeHumidityDehumidifierThreshold = characteristic.NewRelativeHumidityDehumidifierThreshold()
	svc.AddCharacteristic(svc.RelativeHumidityDehumidifierThreshold.Characteristic)

	svc.RelativeHumidityHumidifierThreshold = characteristic.NewRelativeHumidityHumidifierThreshold()
	svc.AddCharacteristic(svc.RelativeHumidityHumidifierThreshold.Characteristic)

	return &svc
}

//AccessoryHumidifierDehumidifier struct
type AccessoryHumidifierDehumidifier struct {
	*accessory.Accessory
	HumidifierDehumidifier *ServiceHumidifierDehumidifier
}

//NewAccessoryHumidifierDehumidifier returns AccessoryHumidifierDehumidifier
//
//args[0](int) - TargetHumidifierDehumidifierState.SetValue(args[0]) default(0)
//
//args[1](int) - TargetHumidifierDehumidifierState.SetMinValue(args[1]) default(0)
//
//args[2](int) - TargetHumidifierDehumidifierState.SetMaxValue(args[2]) default(2)
//
//args[3](int) - TargetHumidifierDehumidifierState.SetStepValue(args[3]) default(1)
func NewAccessoryHumidifierDehumidifier(info accessory.Info, args ...interface{}) *AccessoryHumidifierDehumidifier {
	acc := AccessoryHumidifierDehumidifier{}
	acc.Accessory = accessory.New(info, accessory.TypeDehumidifier)
	acc.HumidifierDehumidifier = NewServiceHumidifierDehumidifier()
	amountArgs := len(args)
	if amountArgs > 0 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetValue(argToInt(args[0], 0))
	} else {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetValue(0)
	}
	if amountArgs > 1 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMinValue(argToInt(args[1], 0))
	} else {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMinValue(0)
	}
	if amountArgs > 2 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMaxValue(argToInt(args[2], 2))
	} else {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMaxValue(2)
	}
	if amountArgs > 3 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetStepValue(argToInt(args[3], 1))
	} else {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetStepValue(1)
	}
	acc.AddService(acc.HumidifierDehumidifier.Service)
	return &acc
}
