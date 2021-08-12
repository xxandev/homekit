package homekit

import (
	haps "github.com/alpr777/homekit/hap-service"
	"github.com/brutella/hc/accessory"
)

//AccessoryDoor struct
type AccessoryDoor struct {
	*accessory.Accessory
	Door *haps.Door
}

//NewAccessoryDoor returns AccessoryDoor
//  args[0](int) - TargetPosition.SetValue(args[0]) default(0)
//  args[1](int) - TargetPosition.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetPosition.SetMaxValue(args[2]) default(2)
//  args[3](int) - TargetPosition.SetStepValue(args[3]) default(1)
func NewAccessoryDoor(info accessory.Info, args ...interface{}) *AccessoryDoor {
	acc := AccessoryDoor{}
	acc.Accessory = accessory.New(info, accessory.TypeDoor)
	acc.Door = haps.NewDoor()
	n := len(args)
	if n > 0 {
		acc.Door.TargetPosition.SetValue(toInt(args[0], 0))
	}
	if n > 1 {
		acc.Door.TargetPosition.SetMinValue(toInt(args[1], 0))
	}
	if n > 2 {
		acc.Door.TargetPosition.SetMaxValue(toInt(args[2], 100))
	}
	if n > 3 {
		acc.Door.TargetPosition.SetStepValue(toInt(args[3], 1))
	}
	acc.AddService(acc.Door.Service)
	return &acc
}
