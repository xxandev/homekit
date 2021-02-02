package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

type serviceDoor struct {
	*service.Service
	CurrentPosition     *characteristic.CurrentPosition
	PositionState       *characteristic.PositionState
	TargetPosition      *characteristic.TargetPosition
	ObstructionDetected *characteristic.ObstructionDetected
}

func newServiceDoor() *serviceDoor {
	svc := serviceDoor{}
	svc.Service = service.New(service.TypeDoor)

	svc.CurrentPosition = characteristic.NewCurrentPosition()
	svc.AddCharacteristic(svc.CurrentPosition.Characteristic)

	svc.PositionState = characteristic.NewPositionState()
	svc.AddCharacteristic(svc.PositionState.Characteristic)

	svc.TargetPosition = characteristic.NewTargetPosition()
	svc.AddCharacteristic(svc.TargetPosition.Characteristic)

	svc.ObstructionDetected = characteristic.NewObstructionDetected()
	svc.AddCharacteristic(svc.ObstructionDetected.Characteristic)

	return &svc
}

//AccessoryDoor struct
type AccessoryDoor struct {
	*accessory.Accessory
	Door *serviceDoor
}

//NewAccessoryDoor returns AccessoryDoor
//
//args[0](int) - TargetPosition.SetValue(args[0]) default(0)
//
//args[1](int) - TargetPosition.SetMinValue(args[1]) default(0)
//
//args[2](int) - TargetPosition.SetMaxValue(args[2]) default(2)
//
//args[3](int) - TargetPosition.SetStepValue(args[3]) default(1)
func NewAccessoryDoor(info accessory.Info, args ...interface{}) *AccessoryDoor {
	acc := AccessoryDoor{}
	acc.Accessory = accessory.New(info, accessory.TypeDoor)
	acc.Door = newServiceDoor()
	amountArgs := len(args)
	if amountArgs > 0 {
		acc.Door.TargetPosition.SetValue(argToInt(args[0], 0))
	} else {
		acc.Door.TargetPosition.SetValue(0)
	}
	if amountArgs > 1 {
		acc.Door.TargetPosition.SetMinValue(argToInt(args[1], 0))
	} else {
		acc.Door.TargetPosition.SetMinValue(0)
	}
	if amountArgs > 2 {
		acc.Door.TargetPosition.SetMaxValue(argToInt(args[2], 100))
	} else {
		acc.Door.TargetPosition.SetMaxValue(100)
	}
	if amountArgs > 3 {
		acc.Door.TargetPosition.SetStepValue(argToInt(args[3], 1))
	} else {
		acc.Door.TargetPosition.SetStepValue(1)
	}
	acc.AddService(acc.Door.Service)
	return &acc
}
