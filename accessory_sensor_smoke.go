package homekit

import (
	"fmt"
	"time"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorSmoke struct
type AccessorySensorSmoke struct {
	*accessory.A
	SmokeSensor *service.SmokeSensor
}

func (acc *AccessorySensorSmoke) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorSmoke) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorSmoke) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorSmoke) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorSmoke) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorSmoke) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorSmoke return AccessorySensorSmoke (args... are not used)
func NewAccessorySensorSmoke(info accessory.Info, args ...interface{}) *AccessorySensorSmoke {
	acc := AccessorySensorSmoke{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.SmokeSensor = service.NewSmokeSensor()
	acc.AddS(acc.SmokeSensor.S)
	return &acc
}

func (acc *AccessorySensorSmoke) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorSmoke) OnExample() {
	go func() {
		for range time.Tick(10 * time.Second) {
			if acc.SmokeSensor.SmokeDetected.Value() > 0 {
				acc.SmokeSensor.SmokeDetected.SetValue(0)
			} else {
				acc.SmokeSensor.SmokeDetected.SetValue(1)
			}
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update current state: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.SmokeSensor.SmokeDetected.Value())
		}
	}()
}
