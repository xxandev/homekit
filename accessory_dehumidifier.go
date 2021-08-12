package homekit

import (
	haps "github.com/alpr777/homekit/hap-service"
	"github.com/brutella/hc/accessory"
)

//AccessoryHumidifierDehumidifier struct
type AccessoryHumidifierDehumidifier struct {
	*accessory.Accessory
	HumidifierDehumidifier *haps.HumidifierDehumidifier
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
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetValue(toInt(args[0], 0))
	}
	if n > 1 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMinValue(toInt(args[1], 0))
	}
	if n > 2 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetMaxValue(toInt(args[2], 2))
	}
	if n > 3 {
		acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.SetStepValue(toInt(args[3], 1))
	}
	acc.AddService(acc.HumidifierDehumidifier.Service)
	return &acc
}
