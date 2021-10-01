package homekit

import (
	haps "github.com/alpr777/homekit/hap-service"
	"github.com/brutella/hc/accessory"
)

//AccessoryFanControlled struct
type AccessoryFanControlled struct {
	*accessory.Accessory
	Fan *haps.FanControlled
}

//NewAccessoryFanControlled return AccessoryFanControlled
//  args[0](float64) - RotationSpeed.SetValue(args[0]) default(0)
//  args[1](float64) - RotationSpeed.SetMinValue(args[1]) default(0)
//  args[2](float64) - RotationSpeed.SetMaxValue(args[2]) default(100)
//  args[3](float64) - RotationSpeed.SetStepValue(args[3]) default(1)
func NewAccessoryFanControlled(info accessory.Info, args ...interface{}) *AccessoryFanControlled {
	acc := AccessoryFanControlled{}
	acc.Accessory = accessory.New(info, accessory.TypeFan)
	acc.Fan = haps.NewFanControlled()
	n := len(args)
	if n > 0 {
		acc.Fan.RotationSpeed.SetValue(toFloat64(args[0], 0.0))
	}
	if n > 1 {
		acc.Fan.RotationSpeed.SetMinValue(toFloat64(args[1], 0.0))
	}
	if n > 2 {
		acc.Fan.RotationSpeed.SetMaxValue(toFloat64(args[2], 100.0))
	}
	if n > 3 {
		acc.Fan.RotationSpeed.SetStepValue(toFloat64(args[3], 1.0))
	}
	acc.AddService(acc.Fan.Service)
	return &acc
}

func (acc *AccessoryFanControlled) OnValuesRemoteUpdateEmpty(fn func()) {
	acc.Fan.On.OnValueRemoteUpdate(func(_ bool) { fn() })
	acc.Fan.RotationSpeed.OnValueRemoteUpdate(func(_ float64) { fn() })
}
