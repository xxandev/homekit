package homekit

import (
	"github.com/brutella/hap/accessory"
	haps "github.com/xxandev/homekit/hap-service"
)

//AccessoryHumidifierDehumidifier struct
type AccessoryHumidifierDehumidifier struct {
	*accessory.A
	HumidifierDehumidifier *haps.HumidifierDehumidifier
}

func (acc *AccessoryHumidifierDehumidifier) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryHumidifierDehumidifier) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryHumidifierDehumidifier) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryHumidifierDehumidifier) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryHumidifierDehumidifier) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryHumidifierDehumidifier) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryHumidifierDehumidifier returns *HumidifierDehumidifier
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args[0](int) - TargetHumidifierDehumidifierState.SetValue(args[0]) default(0)
//  args[1](int) - TargetHumidifierDehumidifierState.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetHumidifierDehumidifierState.SetMaxValue(args[2]) default(2)
//  args[3](int) - TargetHumidifierDehumidifierState.SetStepValue(args[3]) default(1)
func NewAccessoryHumidifierDehumidifier(info accessory.Info, args ...interface{}) *AccessoryHumidifierDehumidifier {
	acc := AccessoryHumidifierDehumidifier{}
	acc.A = accessory.New(info, accessory.TypeDehumidifier)
	acc.HumidifierDehumidifier = haps.NewHumidifierDehumidifier()
	n := len(args)
	if n > 0 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetValue(toi(args[0], 0))
	}
	if n > 1 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMinValue(toi(args[1], 0))
	}
	if n > 2 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMaxValue(toi(args[2], 2))
	}
	if n > 3 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetStepValue(toi(args[3], 1))
	}
	acc.AddS(acc.HumidifierDehumidifier.S)
	return &acc
}

//NewAccessoryHumidifierDehumidifier returns *HumidifierDehumidifier
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args[0](int) - TargetHumidifierDehumidifierState.SetValue(args[0]) default(0)
//  args[1](int) - TargetHumidifierDehumidifierState.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetHumidifierDehumidifierState.SetMaxValue(args[2]) default(2)
//  args[3](int) - TargetHumidifierDehumidifierState.SetStepValue(args[3]) default(1)
func NewAccHumidifierDehumidifier(id uint64, info accessory.Info, args ...interface{}) *AccessoryHumidifierDehumidifier {
	acc := AccessoryHumidifierDehumidifier{}
	acc.A = accessory.New(info, accessory.TypeDehumidifier)
	acc.HumidifierDehumidifier = haps.NewHumidifierDehumidifier()
	n := len(args)
	if n > 0 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetValue(toi(args[0], 0))
	}
	if n > 1 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMinValue(toi(args[1], 0))
	}
	if n > 2 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMaxValue(toi(args[2], 2))
	}
	if n > 3 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetStepValue(toi(args[3], 1))
	}
	acc.AddS(acc.HumidifierDehumidifier.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessoryHumidifierDehumidifier) OnValuesRemoteUpdates(fn func()) {
	acc.HumidifierDehumidifier.Active.OnValueRemoteUpdate(func(int) { fn() })
	acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.OnValueRemoteUpdate(func(int) { fn() })
	acc.HumidifierDehumidifier.RelativeHumidityDehumidifierThreshold.OnValueRemoteUpdate(func(float64) { fn() })
	acc.HumidifierDehumidifier.RelativeHumidityHumidifierThreshold.OnValueRemoteUpdate(func(float64) { fn() })
}
