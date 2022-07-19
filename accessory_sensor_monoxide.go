package homekit

import (
	"fmt"
	"time"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorMonoxide struct
type AccessorySensorMonoxide struct {
	*accessory.A
	CarbonMonoxideSensor *service.CarbonMonoxideSensor
}

func (acc *AccessorySensorMonoxide) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorMonoxide) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorMonoxide) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorMonoxide) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorMonoxide) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorMonoxide) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorMonoxide return AccessorySensorMonoxide (args... are not used)
func NewAccessorySensorMonoxide(info accessory.Info, args ...interface{}) *AccessorySensorMonoxide {
	acc := AccessorySensorMonoxide{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.CarbonMonoxideSensor = service.NewCarbonMonoxideSensor()
	acc.AddS(acc.CarbonMonoxideSensor.S)
	return &acc
}

func (acc *AccessorySensorMonoxide) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorMonoxide) OnExample() {
	go func() {
		for range time.Tick(10 * time.Second) {
			if acc.CarbonMonoxideSensor.CarbonMonoxideDetected.Value() > 0 {
				acc.CarbonMonoxideSensor.CarbonMonoxideDetected.SetValue(0)
			} else {
				acc.CarbonMonoxideSensor.CarbonMonoxideDetected.SetValue(1)
			}
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update current state: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.CarbonMonoxideSensor.CarbonMonoxideDetected.Value())
		}
	}()
}
