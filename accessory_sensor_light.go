package homekit

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorLight struct
type AccessorySensorLight struct {
	*accessory.A
	LightSensor *service.LightSensor
}

func (acc *AccessorySensorLight) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorLight) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorLight) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorLight) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorLight) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorLight) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorLight return AccessorySensorLight (args... are not used)
func NewAccessorySensorLight(info accessory.Info, args ...interface{}) *AccessorySensorLight {
	acc := AccessorySensorLight{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.LightSensor = service.NewLightSensor()
	acc.AddS(acc.LightSensor.S)
	return &acc
}

func (acc *AccessorySensorLight) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorLight) OnExample() {
	go func() {
		for range time.Tick(10 * time.Second) {
			acc.LightSensor.CurrentAmbientLightLevel.SetValue(rand.Float64())
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update current state: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.LightSensor.CurrentAmbientLightLevel.Value())
		}
	}()
}
