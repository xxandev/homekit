package homekit

import (
	"fmt"
	"time"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorMotion struct
type AccessorySensorMotion struct {
	*accessory.A
	MotionSensor *service.MotionSensor
}

func (acc *AccessorySensorMotion) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorMotion) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorMotion) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorMotion) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorMotion) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorMotion) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorMotion return AccessorySensorMotion (args... are not used)
func NewAccessorySensorMotion(info accessory.Info, args ...interface{}) *AccessorySensorMotion {
	acc := AccessorySensorMotion{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.MotionSensor = service.NewMotionSensor()
	acc.AddS(acc.MotionSensor.S)
	return &acc
}

func (acc *AccessorySensorMotion) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorMotion) OnExample() {
	go func() {
		for range time.Tick(10 * time.Second) {
			acc.MotionSensor.MotionDetected.SetValue(!acc.MotionSensor.MotionDetected.Value())
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update current state: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.MotionSensor.MotionDetected.Value())
		}
	}()
}
