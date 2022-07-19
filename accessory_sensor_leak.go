package homekit

import (
	"fmt"
	"time"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorLeak struct
type AccessorySensorLeak struct {
	*accessory.A
	LeakSensor *service.LeakSensor
}

func (acc *AccessorySensorLeak) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorLeak) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorLeak) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorLeak) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorLeak) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorLeak) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorLeak return AccessorySensorLeak (args... are not used)
func NewAccessorySensorLeak(info accessory.Info, args ...interface{}) *AccessorySensorLeak {
	acc := AccessorySensorLeak{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.LeakSensor = service.NewLeakSensor()
	acc.AddS(acc.LeakSensor.S)
	return &acc
}

func (acc *AccessorySensorLeak) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorLeak) OnExample() {
	go func() {
		for range time.Tick(30 * time.Second) {
			if acc.LeakSensor.LeakDetected.Value() > 0 {
				acc.LeakSensor.LeakDetected.SetValue(0)
			} else {
				acc.LeakSensor.LeakDetected.SetValue(1)
			}
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update current state: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.LeakSensor.LeakDetected.Value())
		}
	}()
}
