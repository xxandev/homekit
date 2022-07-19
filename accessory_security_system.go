package homekit

import (
	"fmt"
	"time"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySecuritySystem struct
type AccessorySecuritySystem struct {
	*accessory.A
	SecuritySystemSimple *service.SecuritySystem
}

func (acc *AccessorySecuritySystem) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySecuritySystem) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySecuritySystem) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySecuritySystem) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySecuritySystem) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySecuritySystem) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySecuritySystem returns AccessorySecuritySystem
//  args[0](int) - SecuritySystemTargetState.SetValue(args[0]) default(0)
//  args[1](int) - SecuritySystemTargetState.SetMinValue(args[1]) default(0)
//  args[2](int) - SecuritySystemTargetState.SetMaxValue(args[2]) default(3)
//  args[3](int) - SecuritySystemTargetState.SetStepValue(args[3]) default(1)
func NewAccessorySecuritySystem(info accessory.Info, args ...interface{}) *AccessorySecuritySystem {
	acc := AccessorySecuritySystem{}
	acc.A = accessory.New(info, accessory.TypeSecuritySystem)
	acc.SecuritySystemSimple = service.NewSecuritySystem()
	n := len(args)
	if n > 0 {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetValue(toi(args[0], 0))
	}
	if n > 1 {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetMinValue(toi(args[1], 0))
	}
	if n > 2 {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetMaxValue(toi(args[2], 3))
	}
	if n > 3 {
		acc.SecuritySystemSimple.SecuritySystemTargetState.SetStepValue(toi(args[3], 1))
	}
	acc.AddS(acc.SecuritySystemSimple.S)
	return &acc
}

func (acc *AccessorySecuritySystem) OnValuesRemoteUpdates(fn func()) {
	acc.SecuritySystemSimple.SecuritySystemTargetState.OnValueRemoteUpdate(func(int) { fn() })
}

func (acc *AccessorySecuritySystem) OnExample() {
	go func() {
		for range time.Tick(30 * time.Second) {
			if acc.SecuritySystemSimple.SecuritySystemCurrentState.Value() >= 4 {
				acc.SecuritySystemSimple.SecuritySystemCurrentState.SetValue(0)
			} else {
				acc.SecuritySystemSimple.SecuritySystemCurrentState.SetValue(acc.SecuritySystemSimple.SecuritySystemCurrentState.Value() + 1)
			}
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update current state: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.SecuritySystemSimple.SecuritySystemCurrentState.Value())
		}
	}()
	acc.SecuritySystemSimple.SecuritySystemTargetState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update target state: %[4]T - %[4]T\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
}
