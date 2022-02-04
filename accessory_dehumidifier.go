package homekit

import (
	"fmt"

	haps "github.com/alpr777/homekit/hap-service"
	"github.com/brutella/hc/accessory"
)

//AccessoryHumidifierDehumidifier struct
type AccessoryHumidifierDehumidifier struct {
	*accessory.Accessory
	HumidifierDehumidifier *haps.HumidifierDehumidifier
}

func (acc *AccessoryHumidifierDehumidifier) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryHumidifierDehumidifier) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryHumidifierDehumidifier) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryHumidifierDehumidifier) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryHumidifierDehumidifier) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryHumidifierDehumidifier returns AccessoryHumidifierDehumidifier
//  args[0](int) - TargetHumidifierDehumidifierState.SetValue(args[0]) default(0)
//  args[1](int) - TargetHumidifierDehumidifierState.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetHumidifierDehumidifierState.SetMaxValue(args[2]) default(2)
//  args[3](int) - TargetHumidifierDehumidifierState.SetStepValue(args[3]) default(1)
func NewAccessoryHumidifierDehumidifier(info accessory.Info, args ...interface{}) *AccessoryHumidifierDehumidifier {
	acc := AccessoryHumidifierDehumidifier{}
	acc.Accessory = accessory.New(info, accessory.TypeDehumidifier)
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
	acc.AddService(acc.HumidifierDehumidifier.Service)
	return &acc
}

func (acc *AccessoryHumidifierDehumidifier) OnValuesRemoteUpdates(fn func()) {
	acc.HumidifierDehumidifier.Active.OnValueRemoteUpdate(func(int) { fn() })
	acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.OnValueRemoteUpdate(func(int) { fn() })
	acc.HumidifierDehumidifier.RelativeHumidityDehumidifierThreshold.OnValueRemoteUpdate(func(float64) { fn() })
	acc.HumidifierDehumidifier.RelativeHumidityHumidifierThreshold.OnValueRemoteUpdate(func(float64) { fn() })
}

func (acc *AccessoryHumidifierDehumidifier) OnValuesRemoteUpdatesPrint() {
	acc.HumidifierDehumidifier.Active.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update active: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update target state: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.HumidifierDehumidifier.RelativeHumidityDehumidifierThreshold.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%T - %s] remote update relative threshold: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.HumidifierDehumidifier.RelativeHumidityHumidifierThreshold.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%T - %s] remote update relative threshold: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
