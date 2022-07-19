package homekit

import (
	"fmt"
	"time"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorContact struct
type AccessorySensorContact struct {
	*accessory.A
	ContactSensor *service.ContactSensor
}

func (acc *AccessorySensorContact) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorContact) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorContact) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorContact) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorContact) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorContact) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorContact return AccessorySensorAirQuality (args... are not used)
func NewAccessorySensorContact(info accessory.Info, args ...interface{}) *AccessorySensorContact {
	acc := AccessorySensorContact{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.ContactSensor = service.NewContactSensor()
	acc.AddS(acc.ContactSensor.S)
	return &acc
}

func (acc *AccessorySensorContact) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorContact) OnExample() {
	go func() {
		for range time.Tick(30 * time.Second) {
			if acc.ContactSensor.ContactSensorState.Value() > 0 {
				acc.ContactSensor.ContactSensorState.SetValue(0)
			} else {
				acc.ContactSensor.ContactSensorState.SetValue(1)
			}
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update current state: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.ContactSensor.ContactSensorState.Value())
		}
	}()
}
